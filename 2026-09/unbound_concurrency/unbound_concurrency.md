---
marp: true
theme: rose-pine
paginate: true
title: "Unbound Concurrency in Go"
---

<style>
code { font-size: 24px; }
section { justify-content: flex-start; }
</style>

# Unbound Concurrency in Go

### A guided tour of a port scanner, evolved

- Based on the GopherCon 2021 talk by **Johnny Boursiquot** & **Aaron Schlesinger**
- Code re-implemented & updated by **Peter Preeper**
- One small program, ~10 evolutions, 3 years of production lessons

<!--
Welcome everyone. Today we're going to look at concurrency in Go through the lens of a single program: a port scanner. The original talk is from GopherCon 2021 by two senior engineers, Johnny Boursiquot and Aaron Schlesinger. I've taken the code they showed and updated it to Go 1.25 idioms (you'll notice the new WaitGroup.Go method in a few of the later versions), so what you'll see here runs today.

The whole arc: start with the simplest correct thing, then keep evolving it as we learn what the runtime needs from us. Nothing here is exotic — it's almost entirely the standard library.
-->

---

# What is a port scan?

- Ask a host "are you listening on this TCP port?"
- **Open** → the connect succeeded
- **Closed** → the connect failed. Any error. Refused, timeout, unreachable…
- Every port gets the *same, independent* probe

<!--
A port scan is conceptually simple: try to open a TCP connection to host:port. If the connection succeeds, the port is open — a process is listening. If it fails for any reason, we call it closed. We don't care about the specific error; the probe is binary.

The key observation for today: every probe is identical and independent. Port 80's probe does not depend on port 443's probe.
-->

---

# The probe: the TCP handshake

```
SYN ─────────────▶   ♪
SYN-ACK ◀─────────   ♪   → dialing the port
ACK ─────────────▶   ♪   → connection established = OPEN
```

- `net.Dial("tcp", net.JoinHostPort(host, port))` is a **connection attempt**
- Wrap up the handshake, return a connection — or an error

<!--
All `net.Dial` does (all it can do) is engage that three-way handshake under the hood. It's the network package's abstraction: give it host and port, it returns either a live connection or an error.

If we get a connection back, a process was bound and listening. If we get an error, nothing was listening. That's the whole scanner primitive. One line of the standard library.
-->

---

# The philosophy: start simple

> "Start with the simplest thing that could possibly work."

- ~90% of what follows is **standard library**
- One third-party package: `golang.org/x/sync` (the semaphore)
- ~10 evolutions ≈ 3 years of production experience, compressed into one talk

<!--
The speakers set the ground rule immediately: don't design a concurrent system upfront. Write the simplest thing that could possibly work, then evolve it as you hit real constraints. Each version of the scanner is a response to a concrete problem the previous version hit.

Also worth noting: nearly everything we use is in the standard library. The only external package in the entire journey is the experimental-but-production-ready semaphore from the x/sync group, which landed in the language's ecosystem around that time. In my updated code it's pinned as a dependency in go.mod.
-->

---

# v1 — `01-sequential`: the simplest thing

<style>
code { font-size: 22px; }
</style>

```go
func main() {
	flag.Parse()
	portsToScan, err := parsePortsToScan(ports)
	...

	for _, p := range portsToScan {
		conn, err := net.Dial("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", p)))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", p, err)
			continue
		}
		conn.Close()
		log.Printf("%d OPEN\n", p)
	}
}
```

`openssl s_client`, but with a loop. ✅ correct, ✅ readable, ❌ slow

<!--
This is version one: a for loop calling net.Dial on each port and printing the result. It's correct, it's readable, and it's slow. One dial at a time, waiting on the network round-trip for every single port.

Note the flags up top — every version of this program shares the same CLI: --host, --ports ("80,22-100"), --timeout, --workers, --outfile. The parse logic parsePortsToScan handles the comma-and-dash syntax and is byte-for-byte identical in all 13 source directories of the repo.

The lesson of v1 is not the code — it's the default to reach for when you don't have a reason to be clever.
-->

---

# But wait — is this parallelizable?

This is **embarrassingly parallel** work:

- No scan shares data or state with another scan
- Each port = one independent unit of work
- We have multiple cores… and they sit idle while one dial at a time happens

The **only** constraint: every goroutine must eventually finish and report.

<!--
Port scanning is what we call embarrassingly parallel: there's no dependency between units of work. Nothing produced by the port-80 probe is needed by the port-443 probe. (The speakers fumble the phrase on stage — "embarrassingly parallel" — and it's a running joke early in the talk.)

If your machine has four cores, you'd like four probes actually happening at once. That's the difference between concurrency — lots of units of work — and parallelism — actually using multiple cores. Go gives us both, cheaply.

But the speakers' Rule #1 comes first…
-->

---

# Rule #1: know when goroutines are *done*

> The `go` keyword *schedules* work. It does **not** guarantee the work starts, and it does **not** guarantee it completes.

- Goroutines start running once the scheduler gets to them
- They finish when **they** return — `go` gives you no receipt

<!--
The single most important sentence of the whole talk: `go` only schedules. You hand a function to the scheduler and get back nothing. There is no join, no promise, no future.

So concurrency code fails in a very specific way: main reaches its last line and exits, leaving scheduled-but-not-run goroutines behind, either never started or killed mid-flight. Rule #1 is: always know when your goroutines are actually done — not merely scheduled. Keep that in mind; almost every evolution today is really about enforcing this rule.
-->

---

# v2 — `02-toofast`: the `go` trap

<style>
code { font-size: 22px; }
</style>

```go
for _, p := range portsToScan {
	go func(p int) {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", p))
		if err != nil {
			log.Printf("%d CLOSED (%s)\n", p, err)
			return
		}
		conn.Close()
		log.Printf("%d OPEN\n", p)
	}(p)
}
log.Println("DONE")
```

**Output:** `DONE` — and nothing else. Zero scans.

<!--
So: same loop, same dial, but each probe is now wrapped in `go`. We schedule a goroutine per port and then… print DONE immediately. The scheduler hasn't run a single one of our goroutines by the time main hits its last line.

Two lessons here. First, Rule #1 in action: this program has no idea when its work is done, because it never asked. Second — do you spot the other bug? The dial uses `localhost:%d` instead of the host flag. My update of the talk's code fixed this in later versions so the host flag is threaded through. But for now: scheduling is not execution, and this is the fastest program in the world that accomplishes nothing.
-->

---

# v3 — `03-synchronized`: WaitGroup

<style>
code { font-size: 22px; }
</style>

```go
var wg sync.WaitGroup
numGoRoutinesToWaitOn := len(portsToScan)
wg.Add(numGoRoutinesToWaitOn)

for _, p := range portsToScan {
	go func(p int) {
		defer wg.Done()          // ← decrement when THIS goroutine returns
		conn, err := net.Dial("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", p)))
		...
	}(p)
}

wg.Wait()
log.Println("DONE")
```

- `Add(n)` **once**, *before* launching → `Done()` n times → `Wait()` returns
- Blocking: the scheduler pauses this goroutine and hands control back when the counter hits 0

<!--
WaitGroup is Go's answer to Rule #1. It's a counter with a waiting room: Add increments, Done decrements, Wait blocks until the counter reaches zero.

The zero value is the useful part — you declare a WaitGroup and use it, no initialization needed. And the shape matters: Add the total ONCE before you launch anything. The anti-pattern the speakers warn about is calling Add inside the goroutine, right next to defer Done — it's "smelly" code that works by accident of timing.

How a sync-less walkthrough reads: main runs into wg.Wait, the scheduler pauses main's thread, runs our goroutines, each one's deferred Done decrements, and when the counter hits zero the scheduler hands control back to main. Program is correct again — and one core, still.
-->

---

# v3.5 — the Go 1.25 update: `WaitGroup.Go`

<style>
code { font-size: 22px; }
</style>

```go
var wg sync.WaitGroup
// no sizing the waitgroup
for _, p := range portsToScan {
	// wg.Go added in go1.25
	wg.Go(func() {
		// no defer wg.Done()
		conn, err := net.Dial("tcp", net.JoinHostPort(host, fmt.Sprintf("%d", p)))
		...
	})
}

wg.Wait()
log.Println("DONE")
```

- `wg.Go(func)` = `Add(1)` + `go` + `defer Done()` in one
- No sizing, no manual bookkeeping — the group counts itself

<!--
This is the version you get in my updated repo. Since Go 1.25, WaitGroup has a Go method that wraps the whole Add/go/defer-Done dance into a single call. Notice what disappears: the sizing line, the Add call, and the deferred Done inside the closure. The code comments in the repo call out the 1.25 addition explicitly.

The lesson: the standard library keeps absorbing the patterns the talk taught. What took two moving parts in 2021 is one method today. But don't run this yet — running it properly introduces our next bug, on purpose.
-->

---

# Unbounded concurrency: the blow-up

- One goroutine per port × **5,000–65,000** ports
- Every dial holds a **file descriptor** (FD)
- OS error you will see: **"too many open files"** ✋

```
portscan: error: dial tcp 127.0.0.1:63521: connect: too many open files
```

- This is the **OS protecting *itself*** — not your bug. Your bug is asking for it.

<!--
Now run the synchronized version against a wide range — say 5,000 to 65,000 ports. We schedule potentially tens of thousands of goroutines, each holding a socket, each socket a file descriptor. The operating system has a hard per-process limit on open file descriptors, and when we blow through it, every dial starts failing with "too many open files".

This looks like a crash, but read it the right way: the OS refusing to hand out more FDs is the OS protecting itself and everyone else on the box from our greed. The bug is ours — we asked for unlimited concurrency. The speakers' guidance: find and work within the constraints of your environment. File descriptors, network throughput, database connections, latency budgets — understand which one you're bound by before you pick a design.
-->

---

# v4 — `04-workerpool`: bound it

<style>
code { font-size: 22px; }
</style>

```go
portsChan := make(chan int, workers)   // buffered chan = a queue
resultsChan := make(chan int)

for i := 0; i < cap(portsChan); i++ { // numWorkers also acceptable here
	go worker(host, portsChan, resultsChan)
}
```

- A **fixed** N goroutines (default `NumCPU`) consume work off a queue
- Two channels: work in (`portsChan`), results out (`resultsChan`)

<!--
The fix: stop creating a goroutine per port. Create exactly N worker goroutines — the default is the number of logical CPUs — and hand them work through channels.

Work flows in on portsChan, results come back on resultsChan. The worker count is keyed off cap(portsChan) — the size of the buffered queue — and the repo comment notes numWorkers would be equally acceptable, since both derive from the same --workers flag.
-->

---

# v4, continued — feed & drain

<style>
code { font-size: 22px; }
</style>

```go
go func() {
	for _, p := range portsToScan {
		portsChan <- p                 // producer runs in the background
	}
}()

for range portsToScan {
	if p := <-resultsChan; p != 0 {   // non-zero port means it's open
		openPorts = append(openPorts, p)
	}
}

close(portsChan)
close(resultsChan)
```

- Sending on a **full** channel blocks → natural back-pressure
- The producer must be its own goroutine, or main deadlocks

<!--
A channel with a capacity is a bounded queue: the buffer holds at most workers items, and when it's full the sender's channel-send blocks. That blocking IS the throttle — we never get more work in flight than the pipeline can handle, so we can never blow past the FD limit again. It operates like a conveyor belt: it only moves as fast as the slowest station.

And the producer must run in its own goroutine because main is busy draining resultsChan. If main had to feed portsChan itself, a full buffer would block main while it's also waiting on results — and both sides would be stuck waiting on each other. Decoupled producer and consumer is what keeps the pipeline moving.
-->

---

# The worker: directional channels

<style>
code { font-size: 22px; }
</style>

```go
func worker(host string, portsChan <-chan int, resultsChan chan<- int) {
	for p := range portsChan {
		address := net.JoinHostPort(host, fmt.Sprintf("%d", p))
		conn, err := net.Dial("tcp", address)
		if err != nil {
			fmt.Printf("%d CLOSED (%s)\n", p, err)
			resultsChan <- 0
			continue
		}
		conn.Close()
		resultsChan <- p
	}
}
```

- `<-chan int` = **read-only** (you can only receive)
- `chan<- int` = **write-only** (you can only send)
- The *compiler* enforces direction → a whole class of tests, for free

<!--
Two details that make the worker pool robust.

First, directionality. The worker's parameters say it all: it reads ports off portsChan and writes results onto resultsChan. `<-chan` is a read-only channel, `chan<-` is a write-only channel, and the compiler refuses to let you do otherwise. If anyone accidentally sends on portsChan or receives on resultsChan, it doesn't compile. The speakers call this out as the happy by-product: a whole class of tests you never have to write.

Second, why does the producer need to be its own goroutine? Because main is busy draining resultsChan with the for-range below. If main had to feed portsChan itself, a full buffer would block main — which is also draining results — deadlock. Producer and consumer must both be independent goroutines, or they can each be blocked waiting for the other. That n+1 pattern — workers plus a feeding goroutine — is the shape of the whole program.
-->

---

# v4, continued: results & signals

<style>
code { font-size: 22px; }
</style>

```go
go func() {                    // SIGINT / SIGTERM → print & exit cleanly
	<-sigs
	printResults(openPorts)
	os.Exit(0)
}()

for range portsToScan {        // main drains results; order is NOT guaranteed
	if p := <-resultsChan; p != 0 {
		openPorts = append(openPorts, p)
	}
}

func printResults(ports []int) {
	sort.Ints(ports)           // sort → deterministic output
	...
}
```

- Results arrive **out of order** → that's *parallelism* showing itself
- A channel + callback handles Ctrl-C so you always leave with clean output

<!--
Two finishing touches in v4.

The signal handler: a channel of os.Signal, a signal.Notify subscription for SIGINT and SIGTERM, and one goroutine that just waits on the channel. When you hit Ctrl-C, the signal arrives as a value on the channel, that goroutine wakes, prints whatever results have landed so far, and exits cleanly. In a concurrent program you can't always guarantee a graceful exit path — so make one explicitly.

And the draining loop: ports come back out of resultsChan in no particular order. That's not a bug — that's parallelism made visible. Workers finish probes at different times. So the final printResults sorts before printing, which keeps the reports deterministic. The wide range is now safe to run.
-->

---

# v5 — `05-semaphore`: bound it, batch-style

<style>
code { font-size: 22px; }
</style>

```go
var semMaxWeight int64 = 100_000
var semAcquisitionWeight int64 = 100

sem := semaphore.NewWeighted(semMaxWeight)
ctx := context.Background()

for _, port := range portsToScan {
	if err := sem.Acquire(ctx, semAcquisitionWeight); err != nil { break }

	go func(port int) {
		defer sem.Release(semAcquisitionWeight)
		p := scan(host, port)
		if p != 0 { openPorts = append(openPorts, p) }
	}(port)
}

// We block here until done.
sem.Acquire(ctx, int64(semMaxWeight))
```

**Restaurant:** the venue allows 100,000 kg total. Your party needs a 100 kg table.

<!--
Worker pool = streaming bound: items flow through at a fixed rate. Semaphore = batch bound: launch everything, but only as many goroutines as the semaphore admits at once.

Think of a restaurant. The room has a total weight capacity — that's maxWeight, here 100,000. Your party needs a certain number of seats — that's acquisitionWeight, here 100. Acquire is standing in the lobby waiting for a table. As parties leave, they Release their weight, and the lobby empties into the dining room. So at most 1,000 parties of 100 can be seated from our 100,000 capacity.

The interesting ending: to know when all batches are done, we try to Acquire the ENTIRE maximum weight. That can only succeed when the semaphore is empty — every party has left. The last Acquire is simultaneously our WaitGroup.

Note: this is the one piece that isn't plain stdlib — the weighted semaphore comes from golang.org/x/sync, the experimental-but-production-ready package pinned in our go.mod.
-->

---

# v6 & v7 — timeouts

**v6 — whole-run timeout:**
```go
ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
defer cancel()
```

**v7 — per-batch timeout** *(acquisitionWeight = 1, maxWeight = NumCPU):*
```go
go func(port int) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	defer cancel()
	if err := sem.Acquire(ctx, semAcquisitionWeight); err != nil { return }
	defer sem.Release(semAcquisitionWeight)
	...
}(port)
```

<!--
Two quick variations, one new concept: a cancellable context.

v6 applies a timeout to the entire run. The deadline starts when we create the context, and both Acquire calls run against it — if the run takes longer than our --timeout flag, acquires start failing and we break out.

The speakers originally wrote it this way, then realized the whole-run timeout isn't what you want if you might scan thousands of ports with a deadline measured in seconds — the very first batch eats the budget. So v7 moves the timeout inside: every goroutine gets its own context derived from the parent, with its own deadline, deferring cancel so one timeout can't poison the others. Each batch gets a full timeout; the run is only as slow as slow batches allow.

Real talk: the per-scan-latency path is a candidate for improvement — a dial that hangs can still stall its goroutine. But the semaphore is still holding its maximum weight, so the outer Acquire-all still knows the run isn't done.
-->

---

# v8 — `08-pipeline`: compose the stages

<style>
code { font-size: 24px; }
</style>

```go
// pipeline
scanChan := store(dest, filter(scan(gen(host, portsToScan...))))

// broken up for explainability
// scanChan = gen(portsToScan...)
// scanChan = scan(scanChan)
// scanChan = filter(scanChan)
// scanChan = store(dest, scanChan)
```

Each stage is a function: **`chan in → work → chan out`**

<!--
New era of the program: instead of one script, a PIPELINE of composable stages. Each stage is just a function that takes a channel and returns a channel — and runs its work in its own goroutine. gen produces the work items, scan does the dialing, filter keeps only open ports, store writes CSV.

My update fixed a real bug from the original talk code here: gen originally didn't take the host, so scanOp had no host field and the CSV had no host column. Now gen carries the host into every scanOp, and asSlice/csvHeaders include it. A scan result is now self-describing: host, port, open/closed, the scan error, and the duration.
-->

---

# v8, continued — a stage, up close

<style>
code { font-size: 24px; }
</style>

```go
func scan(in <-chan scanOp) <-chan scanOp {
	out := make(chan scanOp)
	go func() {
		defer close(out)                 // close = "no more data coming"
		for scan := range in {
			...
			out <- scan
		}
	}()
	return out
}
```

- Create the output channel → schedule the goroutine → **return immediately**
- `defer close` inside: closing a channel broadcasts "no more data"

<!--
This is the structural pattern — and you'll see it in every stage: create the output channel, schedule the goroutine, return the channel immediately, and inside the goroutine defer close.

Closing a channel is how a stage broadcasts "no more data" to whoever reads next; the for-range loop on the far side just finishes. The stage is self-contained, and yes — the one-glance takeaway is that a pipeline stage is *declare, launch, hand back*.
-->

---

# The pipeline's dirty secret

**Conveyor belt:** each port moves through gen → scan → filter → store, one stage at a time.

- A stage is idle while the *next* stage works
- The pipeline is correct, readable, composable… and **serialized**

How do we get the parallelism back?

**→ Fan-out / fan-in:**

- Feed the same input to **many** workers (`fan-out`)
- Collect their outputs back into one stream (`fan-in`)

<!--
Here's the catch the talk spends a while on. A pipeline hands each item to exactly one goroutine at a time. The scan stage — our slow stage — processes one port at a time, waiting on network latency for each. So with a single scan stage we've reproduced v1's slowness, just with nicer plumbing. The conveyor belt abandoned the concurrency we gained in v4.

The fix is the classic pattern: fan-out and fan-in. Run MULTIPLE scan goroutines over the same input channel, splitting the work, then merge their outputs back into a single stream before filtering. That is exactly what v9 does.
-->

---

# v9 — `09-fan-out-fan-in`: the fan-out

<style>
code { font-size: 24px; }
</style>

```go
in := gen(host, portsToScan...)   // note: synchronous gen here

// fan-out
sc1 := scan(in)
sc2 := scan(in)
sc3 := scan(in)

for s := range filter(merge(sc1, sc2, sc3)) {
	fmt.Printf("%#v\n", s)
}
```

- **One** input channel, **many** `scan` consumers → work is split across cores
- `merge` collects it all back into a single stream on the far side

<!--
Three scan goroutines now read from the same input channel — that's the fan-out. Each one dials and forwards results, and the network latency of the three probes overlaps on different cores.

Note the synchronous gen here: with a small range it's fine to fill and close the channel before returning — v10 will move the feeding into its own goroutine. The important move is that the scanners are no longer chained single-file; they run beside each other.
-->

---

# v9, continued — the fan-in: `merge`

```go
func merge(chans ...<-chan scanOp) <-chan scanOp {
	out := make(chan scanOp)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, sc := range chans {
		go func(sc <-chan scanOp) {
			for scan := range sc { out <- scan }
			wg.Done()
		}(sc)
	}
	go func() { wg.Wait(); close(out) }()
	return out
}
```

- Merge interleaves all inputs into one stream — a general combinator
- It must **not** close until every input is drained → Rule #1 inside the combinator

<!--
Merge is where the WaitGroup earns its keep. It can't close its output channel until every input channel is drained — close early and the downstream reader sees EOF in the middle of the stream.

So: Add once for each input, each copying goroutine Does when its input is exhausted, and a separate goroutine Waits for the count to hit zero, then closes out. All the stages are just function compositions, so you can also swap filter for a different predicate without touching the pipeline — orthogonal, composable building blocks.
-->

---

# v10 — `10-fan-out-fan-in-with-workers`: graceful shutdown

<style>
code { font-size: 22px; }
</style>

```go
done := make(chan struct{})
defer close(done)
in := gen(done, host, portsToScan...)
var chans []<-chan scanOp
for i := 0; i < workers; i++ {
	chans = append(chans, scan(done, in))
}
for s := range filterErr(done, merge(done, chans...)) {
	fmt.Printf("%#v\n", s)
	done <- struct{}{}   // signal: everyone stop, now
	return
}
```

- A shared `done` channel hands every pipeline goroutine a shutdown button.

<!--
The grand finale. Two things are new.

First, our fan-out count now comes from the workers flag — same default of logical CPUs — so the pipeline scales with the machine instead of hardcoding three scanners.

And second, the essential tool for production concurrency: a DONE channel. An empty struct channel, shared by the entire pipeline, threaded through every stage as the first parameter. The idiomatic syntax everywhere inside, as in gen:

select {
case out <- scanOp{...}:
case <-done:
	return
}

Every stage's loop is wrapped in a select: do your work by default, but if a value ever appears on done — send struct{} or just close the channel — this goroutine returns. Nothing is forced; every goroutine chooses to stop at its next schedule point. That matters because goroutines are NOT garbage collected while they linger. The runtime won't reap a goroutine that never returns. "Always know when your goroutine is going to shut down" is Rule #1's production flavor, and select-on-done is how you hand every goroutine a shutdown button.
-->

---

# v10 demo: the OS wins one round

<style>
code { font-size: 22px; }
</style>

```go
func filterErr(done <-chan struct{}, in <-chan scanOp) <-chan scanOp {
	out := make(chan scanOp)
	go func() {
		defer close(out)
		for scan := range in {
			select {
			default:
				if !scan.open && strings.Contains(scan.scanErr, "too many open files") {
					out <- scan
				}
			case <-done:
				return
			}
		}
	}()
	return out
}
```

```
$ portscan -ports 1-10000 -workers 100000
12 - open
too many open files
```

<!--
The payoff scene. Run v10 with an absurd worker count against a wide range — deliberately ask the OS for more FDs than it can give. The scan goroutines start failing, and every failure carries scanErr = "too many open files".

FilterErr is a new stage, a select-on-done filter, that watches the results stream for exactly that error. The moment one flows through, main receives it on the filtered channel, sends one struct{} down the done channel, and returns.

Every single goroutine in the pipeline sees that on its next select — gen stops feeding, workers stop dialing, filters stop forwarding, merge stops copying — and each returns. The goroutines can now be collected. One signal, whole program down, cleanly. That's the "unbound concurrency" story finding its landing point: you can go unbound safely IF you can always pull the plug.
-->

---

# The six takeaways

1. **Start simple.** The simplest thing that could possibly work — then evolve.
2. **Find the embarrassingly parallel work** — no inter-task dependencies.
3. **Avoid runaway goroutines.** Know your constraints: FDs, network, DB, latency.
4. **Goroutines are not OS threads.** The OS won't schedule them, limit them, or clean them up.
5. **Lean on the standard library** — WaitGroup, channels, the sync package.
6. **Use advanced patterns only as needed** — semaphores and done-channels earn their keep when you outgrow the basics.

<!--
The whole talk in six lines — worth landing each one deliberately and slowly.

One: start simple. Every program we'd ship today started as the for loop in v1. Two: the reason this whole problem is tractable is that a port scan has no dependencies between probes — identify that shape in your own workloads and concurrency becomes a mechanical exercise. Three: the machine and OS are your neighbors; know what you're asking of them. Four: goroutines are cooperative-cheap, not cavalry-cheap — the OS isn't going to time-slice or count them for you. Five: almost nothing here was third-party. Six: the semaphore and the done-channel are load-bearing only in the later versions — don't front-load that complexity.

The talk's speakers close by pointing people at the portscan repo — and I've continued that lineage, updating the code to Go 1.25 idioms, which is exactly what you have here.
-->

---

# Resources & Q&A

- **Talk:** *Unbound Concurrency in Go* — GopherCon 2021 · Johnny Boursiquot & Aaron Schlesinger
  https://www.youtube.com/watch?v=PAUjYyBfELk
- **This code:** `https://github.com/ppreeper/portscan`
  - 13 source directories, `01-sequential` → `10-fan-out-fan-in-with-workers`
  - `-new` variants exercise the Go 1.25 `WaitGroup.Go`
- Try: `go run 04-workerpool -ports 22-1000 -host scanme.nmap.org`
- Ask me anything — concurrency, channels, semaphores, or shutting things down

<!--
Point people at the sources. The YouTube talk is the primary one — it's short and the speakers are excellent. The repo is laid out as the numbered versions we just walked, so anyone can diff adjacent directories to see exactly what each evolution added.

A good live demo is to take v4 or v5 against a real host and show the out-of-order results and the deterministic sorted printout. For the brave: run v10 with a big worker count and watch filterErr catch the too-many-open-files case live.

Any questions? Happy to go deeper on channels, the semaphore math, or how to know when your own goroutines can safely stop.

— Peter Preeper, Thinksoft Inc.
-->

---

# Thank you!

### Questions & discussion

<!--
Wrap up. Encourage people to try the -new variants and compare with git diff between a version and its -new sibling. Thanks for coming, and remember: know when your goroutines are done.
-->

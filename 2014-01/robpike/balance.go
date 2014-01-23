package main

import (
	"container/heap"
	"fmt"
	"math/rand"
	"time"
)

const nRequester = 100
const nWorker = 10

// Simulation of some work: just sleep for a while and report how long.
func op() int {
	n := rand.Int63n(int64(time.Second))
	time.Sleep(time.Duration(nWorker * n))
	return int(n)
}

// The requester sends Requests to the balancer.
type Request struct {
	fn func() int // The operation to perform.
	c  chan int   // The channel to return the result.
}

// An artificial but illustrative simulation of a requester, a load generator.
func requester(work chan<- Request) {
	c := make(chan int)
	for {
		// Kill some time (fake load).
		time.Sleep(time.Duration(rand.Int63n(int64(nWorker * 2 * time.Second))))
		work <- Request{op, c} // send request
		<-c                    // wait for answer
		// furtherProcess(result)
	}
}

// A channel of requests, plus some load tracking data.
type Worker struct {
	i        int          // index in the heap
	requests chan Request // work to do (buffered channel)
	pending  int          // count of pending tasks
}

// Balancer sends request to most lightly loaded worker.
func (w *Worker) work(done chan *Worker) {
	// Could run the loop body as a goroutine for parallelism.
	for {
		req := <-w.requests // get Request from balancer
		req.c <- req.fn()   // call fn and send result (directly to its requester)
		done <- w           // we've finished this request
	}
}

type Pool []*Worker

// Make Pool an implementation of the Heap interface by providing a few methods.
func (p Pool) Len() int { return len(p) }

func (p Pool) Less(i, j int) bool {
	return p[i].pending < p[j].pending
}

func (p *Pool) Swap(i, j int) {
	a := *p
	a[i], a[j] = a[j], a[i]
	a[i].i = i
	a[j].i = j
}

func (p *Pool) Push(x interface{}) {
	a := *p
	n := len(a)
	a = a[0 : n+1]
	w := x.(*Worker)
	a[n] = w
	w.i = n
	*p = a
}

func (p *Pool) Pop() interface{} {
	a := *p
	*p = a[0 : len(a)-1]
	w := a[len(a)-1]
	w.i = -1 // for safety
	return w
}

// The load balancer needs a pool of workers and a single channel
// to which requesters can report task completion.
type Balancer struct {
	pool Pool
	done chan *Worker
	i    int
}

func NewBalancer() *Balancer {
	done := make(chan *Worker, nWorker)
	b := &Balancer{make(Pool, 0, nWorker), done, 0}
	for i := 0; i < nWorker; i++ {
		w := &Worker{requests: make(chan Request, nRequester)}
		heap.Push(&b.pool, w)
		go w.work(b.done)
	}
	return b
}

func (b *Balancer) balance(work chan Request) {
	for {
		select {
		case req := <-work: // received a Request...
			b.dispatch(req) // ...so send it to a Worker
		case w := <-b.done: // a worker has finished...
			b.completed(w) // ...so update its info
		}
		b.print()
	}
}

func (b *Balancer) print() {
	sum := 0
	sumsq := 0
	for _, w := range b.pool {
		fmt.Printf("%d ", w.pending)
		sum += w.pending
		sumsq += w.pending * w.pending
	}
	avg := float64(sum) / float64(len(b.pool))
	variance := float64(sumsq)/float64(len(b.pool)) - avg*avg
	fmt.Printf(" %.2f %.2f\n", avg, variance)
}

// Send Request to worker
func (b *Balancer) dispatch(req Request) {
	w := heap.Pop(&b.pool).(*Worker) // Grab the least loaded worker...
	w.requests <- req                // ...send it the task.
	w.pending++                      // One more in its work queue.
	//	fmt.Printf("started %p; now %d\n", w, w.pending)
	heap.Push(&b.pool, w) // Put it into its place on the heap.
}

// Job is complete; update heap
func (b *Balancer) completed(w *Worker) {
	w.pending-- // One fewer in the queue.
	//	fmt.Printf("finished %p; now %d\n", w, w.pending)
	heap.Remove(&b.pool, w.i) // Remove it from heap.
	heap.Push(&b.pool, w)     // Put it into its place on the heap.
}

func main() {
	work := make(chan Request)
	for i := 0; i < nRequester; i++ {
		go requester(work)
	}
	NewBalancer().balance(work)
}

---
marp: true
theme: default
class:
  - invert
paginate: true
---

<style>
section {
  padding-bottom: 20%;
}
</style>

# Go 1.25 Release

## 🚀 Go 1.25 is here! Performance wins, smarter runtime, and powerful testing tools

Peter Preeper
2025-08-20
ppreeper@gmail.com

---

# Go 1.25

Go 1.25 just dropped (Aug 12, 2025) with a slew of improvements, from container smarts to faster crypto, better debugging, and developer quality-of-life gains, all without breaking Go 1 compatibility.

---

## 🔹 Highlights

---

## 🖥 Smarter for Containers

- On Linux, GOMAXPROCS now adapts to CPU limits in cgroups, perfect for Kubernetes.
- Updates automatically when limits change.

---

## 🧹 Experimental “green tea” GC (opt-in via GOEXPERIMENT=greenteagc)

- More locality & scalability for small-object workloads
- Early tests: 10–40% GC overhead reduction

---

## 📡 Trace Flight Recorder

- Continuous lightweight runtime trace, snapshot anytime
- Great for catching rare prod issues without big overhead

---

## 🐛 Nil-check fix

- Long-standing bug fixed - nil-pointer checks fire in the right place again.

---

## ⚡ Build & Binary Improvements

- DWARF5 debug info by default → smaller debug size & faster links
- Faster slice allocation, better stack use

---

## 🧪 New Testing Superpowers

- testing/synctest for deterministic concurrency tests
- Structured logging via t.Attr
- Cleaner parallel safety with AllocsPerRun

---

## 📦 Experimental JSON v2

- Faster decoder, more config options via GOEXPERIMENT=jsonv2

---

## 🔧 Go Command Enhancements

- -asan with C leak detection by default
- go doc -http opens local docs in browser
- go version -m -json for embedded build info checks

---

## 👥 Concurrency Ergonomics

- sync.WaitGroup.Go for safer goroutine launches

---

## 🔐 Security & Crypto

- TLS 1.2 SHA‑1 signatures disabled by default
- Faster RSA keygen, SHA improvements, and big speedups in FIPS mode

---

## 🌐 Platform Updates

- net/http adds tokenless CSRF defense (CrossOriginProtection)
- Windows async I/O; macOS 12+ now required

---

## 💡 Why it matters

- Container-ready: No more manual tuning for CPUs
- Lower overhead debugging: Trace rare prod bugs with minimal cost
- Performance boosts: From GC, linking, slices, and crypto
- Reliable tests: Bye-bye flaky concurrency test runs

---

## ✅ Pro tip to adopt quickly:

- Remove manual GOMAXPROCS in Kubernetes
- Try new GC in staging for perf wins
- Add CrossOriginProtection to HTTP APIs
- Audit builds using go version -m -json
- Stabilize concurrent tests with testing/synctest

---

## 💬 Which feature are you most excited to try in Go 1.25?

---

## Q&A

- Questions? Let’s discuss!

---

## About Me

- Name: Peter Preeper
- Contact: ppreeper@gmail.com
- I Work For: Thinksoft Inc.
- Job Title: Senior Implementation Specialist

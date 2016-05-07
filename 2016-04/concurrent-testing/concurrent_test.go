package main

import (
	"sync"
	"testing"
	"time"
)

var timeout = time.Millisecond * 10

func blockingFunc() {
	ch := make(chan struct{})
	<-ch
}

func TestBlocking(t *testing.T) {
	ch := make(chan struct{})
	go func() { blockingFunc(); close(ch) }()

	select {
	case <-ch:
	case <-time.After(timeout):
		// t.Error("blockingFunc did not complete within", timeout)
	}
}

func TestOrderedExecution(t *testing.T) {
	outOfOrder := make(chan bool)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() { wg.Wait(); outOfOrder <- true }()
	go func() { outOfOrder <- false; wg.Done() }()

	deadlocked := time.After(timeout)
	select {
	case order := <-outOfOrder:
		if order {
			t.Error("Returned from wg.Wait() before call to wg.Done()")
		}
	case <-deadlocked:
		t.Error("wg.Wait() and wg.Done() deadlocked")
	}
}

func funcA(ch chan struct{}) {
	<-ch
}

func funcB(ch chan struct{}) {
	close(ch)
}

func TestOrderingUsingChannelsNaively(t *testing.T) {
	ch := make(chan struct{})
	outOfOrder := make(chan struct{})

	comm := make(chan struct{})
	go func() { funcA(comm); close(outOfOrder) }()
	go func() { close(ch); funcB(comm) }()

	deadlocked := time.After(timeout)
	select {
	case <-ch:
	case <-outOfOrder:
		t.Error("Expected occasionally: funcA returned before funcB was launched")
	case <-deadlocked:
		t.Error("funcA and funcB deadlocked")
	}
}

func TestOrderOfReturnWillNotWork(t *testing.T) {
	t.Skip()
	outOfOrder := make(chan bool)
	comm := make(chan struct{})
	go func() { funcA(comm); outOfOrder <- false }()
	go func() { funcB(comm); outOfOrder <- true }()

	deadlocked := time.After(timeout)
	select {
	case order := <-outOfOrder:
		if order {
			t.Error("Expected occasionally: funcB returned before funcA")
		}
	case <-deadlocked:
		t.Error("funcA and funcB deadlocked")
	}
}

func TestParallel(t *testing.T) {
	// Setup, fail early if possible. This is not run concurrently.
	t.Parallel()
	// Remainder of test is run concurrently with other parallel tests
}

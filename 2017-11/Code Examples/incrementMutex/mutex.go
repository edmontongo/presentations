package main

import (
	"fmt"
	"sync"
)

func addToInt(A *int) {
	*A++

}

func main() {
	A := 0
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(2)
	go func() {
		for i := 0; i < 10000000; i++ {
			lock.Lock()
			addToInt(&A)
			lock.Unlock()
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 10000000; i++ {
			lock.Lock()
			addToInt(&A)
			lock.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(A)
}

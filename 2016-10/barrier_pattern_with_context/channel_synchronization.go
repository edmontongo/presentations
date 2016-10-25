package main

import (
	"fmt"
	"time"
)

func boundedGenerator(lo, hi int) <-chan int {
	out := make(chan int)
	go func() {
		for i := lo; i < hi; i++ {
			time.Sleep(100 * time.Millisecond)
			out <- 10 * i
		}
		close(out)
	}()
	return out
}

func main() {
	for val := range boundedGenerator(1, 10) {
		fmt.Println(val)
	}
}

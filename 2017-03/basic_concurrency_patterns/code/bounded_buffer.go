package main

import "fmt"

var boundedBuffer = make(chan int, 2) // HL

func Producer() { // HL
	for i := 1; i < 6; i++ {
		fmt.Printf("producer %2d: buffer %d/%d\n", i, len(boundedBuffer), cap(boundedBuffer))
		boundedBuffer <- i // HL
	}
	close(boundedBuffer)
}

func Consumer() { // HL
	for i := range boundedBuffer { // HL
		fmt.Printf("\tconsumer %2d: buffer %d/%d\n", i, len(boundedBuffer), cap(boundedBuffer))
	} // HL
}

func main() {
	go Producer() // HL
	Consumer()    // HL
}

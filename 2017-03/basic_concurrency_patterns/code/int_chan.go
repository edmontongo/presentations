package main

import "fmt"

func main() {
	var intChan chan int // HL

	fmt.Println("An uninitialized channel is:", intChan)
	intChan = make(chan int)

	go func() {
		intChan <- 42 // HL
	}()

	fmt.Println("The answer is:", <-intChan) // HL
}

package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // HL

func greet(who string) {
	fmt.Println("Hello,", who)
	wg.Done() // HL
}

func main() {
	wg.Add(2) // HL

	go greet("wait")
	go greet("group")

	wg.Wait() // HL
}

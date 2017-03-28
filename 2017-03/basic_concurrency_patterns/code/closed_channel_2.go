package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // HL

	close(ch)
	val, ok := <-ch // HL
	fmt.Printf("%v, ok=%t\n", val, ok)
}

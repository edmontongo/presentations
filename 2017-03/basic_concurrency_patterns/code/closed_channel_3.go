package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // HL

	close(ch)
	for val := range ch { // prints nothing // HL
		fmt.Printf("%v, ok=%t\n", val)
	}
}

package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // HL

	close(ch)
	val := <-ch
	fmt.Printf("%v\n", val)
}

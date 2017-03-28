package main

import "fmt"

func countdown(n uint) <-chan uint {
	var ret = make(chan uint) // HL

	go func() { // HL
		for ; n > 0; n-- {
			ret <- n
		}
		close(ret) // HL
	}()
	return ret // HL
}

func main() {
	for n := range countdown(10) {
		fmt.Println(n)
	}
}

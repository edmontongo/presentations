package main

import (
	"fmt"
)

func main() {
	var a []uint
	var done = make(chan struct{}) // HL

	go func() {
		for i := uint(1); i <= 10; i++ {
			a = append(a, i)
		}
		close(done) // HL
	}()

	<-done // HL
	fmt.Println(a)
}

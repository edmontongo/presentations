package main

import (
	"fmt"
	"math"
)

func main() {
	var a []uint
	var done = make(chan struct{}) // HL

	go func() {
		for i := uint(1); i <= 10; i++ {
			a = append(a, uint(int64(float32(math.Log2(float64(uint(1)<<i))))))
		}
		done <- struct{}{} // HL
	}()

	<-done         // HL
	fmt.Println(a) // ?
}

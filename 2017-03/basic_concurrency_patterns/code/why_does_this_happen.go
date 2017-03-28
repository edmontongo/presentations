package main

import (
	"fmt"
	"math"
)

func main() {
	var a []uint

	go func() { // Parallelize a highly complex computation // HL
		for i := uint(1); i <= 10; i++ {
			a = append(a, uint(int64(float32(math.Log2(float64(uint(1)<<i))))))
		}
	}()

	fmt.Println(a) // why? // HL
}

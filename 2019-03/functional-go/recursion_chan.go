package main

import (
	"fmt"
	"time"
)

// START OMIT
func CRecurse(n uint32, result chan uint32) {
	if n == 0 {
		result <- n
		return
	}
	go CRecurse(n-1, result) // HL
}

func main() {
	result := make(chan uint32)
	start := time.Now()
	CRecurse(99999999, result)
	zero := <-result // HL
	fmt.Println(zero)
	fmt.Printf("%+v\n", time.Now().Sub(start))
}

// END OMIT

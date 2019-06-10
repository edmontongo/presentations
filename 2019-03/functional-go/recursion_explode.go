package main

import (
	"fmt"
	"time"
)

// START OMIT
func Recurse(n uint32) uint32 {
	if n == 0 {
		return n
	}
	return Recurse(n - 1)
}

func main() {
	start := time.Now()
	fmt.Println(Recurse(99999999))
	fmt.Printf("%+v\n", time.Now().Sub(start))
}

// END OMIT

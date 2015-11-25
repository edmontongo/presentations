package main

import (
	"fmt"
	"time"
)

func main() {
	// BEGIN OMIT
	var x interface{}

	i := 42
	x = i

	fmt.Printf("1st: %T (%v)\n", x, x)

	s := "hello, world"
	x = s

	fmt.Printf("2nd: %T (%v)\n", x, x)
	// END OMIT
	time.Sleep(100 * time.Millisecond)
}

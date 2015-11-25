package main

import (
	"fmt"
)

func main() {
	// BEGIN OMIT
	i := 42
	fmt.Printf("%v\n", i)

	var w interface{}
	w = i
	// What does w look like? // HL
	// END OMIT
	_ = w
}

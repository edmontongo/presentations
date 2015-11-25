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

	switch x := x.(type) {
	case int:
		fmt.Printf("value: %d\n", x)
	default:
		fmt.Println("unexpected type")
	}
	// END OMIT
	time.Sleep(100 * time.Millisecond)
}

package main

import "fmt"

// START OMIT
func IntGenerator(start func() int) func() int { // HL
	gen := start() - 1
	return func() int {
		gen++      // HL
		return gen // HL
	}
}

func main() {
	intGen := IntGenerator(func() int { return 42 }) // HL
	fmt.Println(intGen())
	fmt.Println(intGen())
	fmt.Println(intGen())
	fmt.Println(intGen())
}

// END OMIT

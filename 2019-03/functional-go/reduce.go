package main

import "fmt"

// START OMIT
type T int
type Sequence []T
type Accumulator func(acc, n T) T

func Reduce(s Sequence, f Accumulator, init T) T { // HL
	for _, n := range s {
		init = f(init, n) // HL
	}
	return init
}

func sum(acc T, n T) T { return acc + n } // HL

func main() {
	s := Sequence{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	fmt.Println(Reduce(s, sum, 0))
}

// END OMIT

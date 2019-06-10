package main

import "fmt"

// START OMIT
type T int
type Sequence []T
type Mapper func(T) T

func Map(s Sequence, f Mapper) Sequence { // HL
	result := make(Sequence, len(s))
	for i, n := range s {
		result[i] = f(n) // HL
	}
	return result
}

func square(n T) T { return n * n } // HL

func main() {
	s := Sequence{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	fmt.Println(Map(s, square))
}

// END OMIT

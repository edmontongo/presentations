package main

import "fmt"

// START OMIT
type T int
type Sequence []T
type Predicate func(T) bool

func Filter(s Sequence, p Predicate) Sequence { // HL
	result := make(Sequence, 0)
	for _, n := range s {
		if p(n) {
			result = append(result, n) // HL
		}
	}
	return result
}

func evens(n T) bool { return n%2 == 0 } // HL

func main() {
	s := Sequence{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(s)
	fmt.Println(Filter(s, evens))
}

// END OMIT

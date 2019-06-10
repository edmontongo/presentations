package main

import "fmt"

type T int
type Sequence []T
type Mapper func(T) T
type Predicate func(T) bool
type Accumulator func(acc T, n T) T

func Map(s Sequence, f Mapper) Sequence {
	result := make(Sequence, len(s))
	for i, n := range s {
		result[i] = f(n) // HL
	}
	return result
}

func square(n T) T { return n * n }

func Filter(s Sequence, p Predicate) Sequence {
	result := make(Sequence, 0)
	for _, n := range s {
		if p(n) {
			result = append(result, n) // HL
		}
	}
	return result
}
func evens(n T) bool { return n%2 == 0 }

func Reduce(s Sequence, f Accumulator, init T) T {
	for _, n := range s {
		init = f(init, n) // HL
	}
	return init
}

func sum(acc T, n T) T { return acc + n }

// START OMIT
func main() {
	s := Sequence{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	sm := Map(s, square)      // HL
	sm = Filter(sm, evens)    // HL
	res := Reduce(sm, sum, 0) // HL
	fmt.Printf("%v\n%v", s, res)
}

// END OMIT

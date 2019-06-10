package main

import "fmt"

type T int
type Sequence []T
type Mapper func(T) T
type Predicate func(T) bool
type Accumulator func(acc T, n T) T

func sum(acc T, n T) T { return acc + n }
func square(n T) T     { return n * n }
func evens(n T) bool   { return n%2 == 0 }

// START1 OMIT
func (s Sequence) Map(f Mapper) Sequence { // HL
	result := make(Sequence, len(s))
	for i, n := range s {
		result[i] = f(n) // HL
	}
	return result
}

func (s Sequence) Filter(p Predicate) Sequence { // HL
	result := make(Sequence, 0)
	for _, n := range s {
		if p(n) {
			result = append(result, n) // HL
		}
	}
	return result
}

func (s Sequence) Reduce(f Accumulator, init T) T { // HL
	for _, n := range s {
		init = f(init, n) // HL
	}
	return init
}

// END1 OMIT

// START2 OMIT
func main() {
	s := Sequence{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := s.Map(square).Filter(evens).Reduce(sum, 0) // HL
	fmt.Printf("%v\n%v", s, res)
}

// END2 OMIT

package main

import "fmt"

type T int
type Sequence []T
type Mapper func(T) T
type Predicate func(T) bool
type Accumulator func(acc T, n T) T

// START OMIT
func main() {
	s := Sequence{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	res := T(0)
	for _, e := range s {
		e *= e
		if e%2 != 0 {
			continue
		}
		res += e
	}
	fmt.Printf("%v\n%v", s, res)
}

// END OMIT

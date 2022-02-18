package main

import "fmt"

// START GTEMPLATE OMIT
type GenericSlice[T any] []T

func (g GenericSlice[T]) Print() {
	for _, v := range g {
		fmt.Println(v)
	}
}

func Print[T any](g GenericSlice[T]) {
	for _, v := range g {
		fmt.Println(v)
	}
}

// END GTEMPLATE OMIT

// START GTEMPCALL OMIT
func typeslice() {
	i := GenericSlice[int]{1, 2, 3}

	i.Print() //1 2 3
	Print(i)  //1 2 3

	f := GenericSlice[float64]{1, 2, 3}

	f.Print() //1 2 3
	Print(f)  //1 2 3

	s := GenericSlice[string]{"1", "2", "3"}

	s.Print() //1 2 3
	Print(s)  //1 2 3
}

// END GTEMPCALL OMIT

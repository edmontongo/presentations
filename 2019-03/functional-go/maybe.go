package main

import (
	"errors"
	"fmt"
)

// START1 OMIT
type T int
type Maybe struct {
	Success T
	Error   error
}

func ForSuccess(result T) Maybe {
	return Maybe{Success: result, Error: nil}
}

func ForError(err error) Maybe {
	return Maybe{Success: -1, Error: err}
}

func (e Maybe) Apply(f func(a T) Maybe) Maybe {
	if e.Error != nil {
		return ForError(e.Error)
	}
	return f(e.Success)
}

// END1 OMIT

// START2 OMIT
func add1(a T) Maybe   { return ForSuccess(a + 1) }
func double(a T) Maybe { return ForSuccess(a * 2) }
func square(a T) Maybe { return ForSuccess(a * a) }
func fail(a T) Maybe   { return ForError(errors.New("Oh Noez!")) }

func main() {
	result := ForSuccess(1).Apply(add1).Apply(double).Apply(square) // HL
	fmt.Printf("%+v\n", result)

	result = ForSuccess(1).Apply(fail).Apply(add1).Apply(double).Apply(square) // HL
	fmt.Printf("%+v\n", result)
}

// END2 OMIT

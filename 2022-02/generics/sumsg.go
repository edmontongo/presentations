package main

// START GENFUNC1 OMIT
// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V { // HL
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// END GENFUNC1 OMIT

// START GENFUNC2 OMIT
type Number interface { // HL
	int64 | float64 // HL
} // HL

// SumNumbers sums the values of map m. It supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](m map[K]V) V { // HL
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

// END GENFUNC2 OMIT

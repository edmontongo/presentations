package main

// START REGFUNC OMIT
// SumInts adds together the values of map m
func SumInts(m map[string]int64) int64 { // HL
	var s int64
	for _, v := range m {
		s += v
	}
	return s
}

// SumFloats adds together the values of map m
func SumFloats(m map[string]float64) float64 { // HL
	var s float64
	for _, v := range m {
		s += v
	}
	return s
}

// END REGFUNC OMIT

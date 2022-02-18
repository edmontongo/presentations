package main

import "fmt"

func main() {

	// START REGFUNC OMIT
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Non-Generic Sums: %v and %v\n",
		SumInts(ints),     // HL
		SumFloats(floats)) // HL

	// END REGFUNC OMIT

	// SumInts(floats)

	// START GENFUNC1 OMIT
	fmt.Printf("Generic Sums: %v and %v\n",
		SumIntsOrFloats[string, int64](ints),     // HL
		SumIntsOrFloats[string, float64](floats)) // HL
	// END GENFUNC1 OMIT

	// Constraints
	// START EQSTD OMIT
	fmt.Println("EqualInt(100,100)", EqualInt(100, 100))
	fmt.Println("EqualString(\"a\",\"a\")", EqualString("a", "a"))
	// END EQSTD OMIT
	// START EQGEN OMIT
	fmt.Println("Equal(\"a\", \"a\")", Equal("a", "a"))
	fmt.Println("Equal(\"a\", \"b\")", Equal("a", "b"))
	// END EQGEN OMIT

	// Non-Generics Sort
	// START BSORT OMIT
	dataI := []int{-2, 45, 0, 11, -9}
	fmt.Println("dataI:", dataI)
	BubbleSortInt(&dataI)
	fmt.Println("dataI:", dataI)

	dataF := []float64{-2, 45, 0, 11, -9}
	fmt.Println("dataF:", dataF)
	// BubbleSortInt(&dataF)
	BubbleSortFloat64(&dataF)
	fmt.Println("dataF:", dataF)
	// END BSORT OMIT
	// Generics Sort
	// START BSORTG OMIT
	dataI = []int{-2, 45, 0, 11, -9}
	fmt.Println("dataI:", dataI)
	BubbleSortAlphaNum(&dataI)
	fmt.Println("dataI:", dataI)

	dataF = []float64{-2, 45, 0, 11, -9}
	fmt.Println("dataF:", dataF)
	BubbleSortAlphaNum(&dataF)
	fmt.Println("dataF:", dataF)

	dataR := []string{"p", "e", "t", "e", "r"}
	fmt.Println("dataR:", dataR)
	BubbleSortAlphaNum(&dataR)
	fmt.Println("dataR:", dataR)
	// END BSORTG OMIT

	// generics definitions

	typeslice()

	// generics interfaces

	typeinterface()

}

package main

import (
	"fmt"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main13)
}

func Main13() error {
	r := -1
	var q int

	// start cut
	m := 924744
	n := 750639
	orig_m, orig_n := m, n

	for {
		q = m / n
		r = m % n
		fmt.Printf("%d = %d × %d + %d\n", m, n, q, r)
		if r == 0 {
			break
		}
		m = n
		n = r
	}
	fmt.Printf("gcd = %d\n", n)
	// end cut
	gcd := n

	fmt.Printf("%d / %d = %f\n", orig_m, gcd, float64(orig_m)/float64(gcd))
	fmt.Printf("%d / %d = %f\n", orig_n, gcd, float64(orig_n)/float64(gcd))

	return nil
}

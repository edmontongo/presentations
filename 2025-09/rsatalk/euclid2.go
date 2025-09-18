package main

import (
	"fmt"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main14)
}

func Main14() error {
	r := -1
	var q int

	ap := 1
	a := 0
	bp := 0
	b := 1

	// start cut
	m := 924744
	n := 750639
	orig_m, orig_n := m, n

	for {
		q = m / n
		r = m % n

		ap, a = a, ap-q*a // HL
		bp, b = b, bp-q*b // HL

		fmt.Printf("%d = %d × %d + %d\n", m, n, q, r)
		if r == 0 {
			break
		}
		fmt.Printf("\t%d = %d × %d + %d × %d\n", a*orig_m+b*orig_n, orig_m, a, orig_n, b)

		m = n
		n = r
	}
	fmt.Printf("gcd = %d\n", n)
	// end cut

	return nil
}

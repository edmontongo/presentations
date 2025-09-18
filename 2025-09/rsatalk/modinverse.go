package main

import (
	"fmt"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main16)
}

type gcdResult struct {
	gcd, a, b, orig_m, orig_n int
}

func (r gcdResult) total() int {
	return r.a*r.orig_m + r.b*r.orig_n
}

func Main16() error {

	gcd := func(m int, n int) gcdResult {
		orig_m, orig_n := m, n

		r := -1
		var q int

		ap := 1
		a := 0
		bp := 0
		b := 1

		for {
			q = m / n
			r = m % n

			ap, a = a, ap-q*a // HL
			bp, b = b, bp-q*b // HL

			//fmt.Printf("%d = %d × %d + %d\n", m, n, q, r)
			if r == 0 {
				break
			}
			//fmt.Printf("\t%d = %d × %d + %d × %d\n", a*orig_m+b*orig_n, orig_m, a, orig_n, b)

			m = n
			n = r
		}
		fmt.Printf("gcd = %d\n", n)
		return gcdResult{
			gcd:    n,
			a:      ap,
			b:      bp,
			orig_m: orig_m,
			orig_n: orig_n,
		}
	}

	// start cut
	m := 7
	n := 15
	g := gcd(m, n)
	if g.gcd != 1 {
		return fmt.Errorf("%d and %d not relatively prime", m, n)
	}
	fmt.Printf("%d × %d + %d x %d = %d\n", g.a, m, g.b, n, g.total())

	minv := (n + g.a%n) % n // translate any negative values into range [0, n)
	fmt.Printf("minv = %d ≡ %d (mod %d)\n", g.a, minv, n)
	fmt.Printf("%d × %d = %d ≡ %d (mod %d)\n", m, minv, m*minv, (m*minv)%n, n)

	// end cut

	return nil
}

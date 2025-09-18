package main

import (
	"fmt"
	"talk/lib"
)

func main() {
	lib.MainWrapper(Main12)
}
func Main12() error {
	// start cut
	s := []byte("pmttwewztl")
	for i := 0; i < 25; i++ {
		for j, _ := range s {
			s[j] = (s[j]-'a'+1)%26 + 'a'
		}
		fmt.Printf("%s\n", s)
	}
	// end cut

	return nil
}

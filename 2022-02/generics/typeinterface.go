package main

import "fmt"

// START TINTF OMIT
type Person interface {
	Work()
}

type codemonkey string

func (c codemonkey) Work() { // HL
	fmt.Println(c, "is working")
}

// specify interface behavior as constraint
func WriteCode[T Person](things []T) { // HL
	for _, v := range things {
		v.Work()
	}
}

func typeinterface() {
	var a, b, c codemonkey
	a = "Jonathan Coulton"
	b = "Linus Torvalds"
	c = "Satoshi Nakamoto"
	WriteCode([]codemonkey{a, b, c})
}

// END TINTF OMIT

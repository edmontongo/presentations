package main

import "fmt"

// START OMIT
func Reverse(str string) string {
	if len(str) == 0 {
		return str
	}
	return Reverse(str[1:]) + str[:1] // HL
}

func main() {
	const str = "0123456789"
	fmt.Println(str)
	fmt.Println(Reverse(str))
}

// END OMIT
//return str[len(str)-1:] + Reverse(str[:len(str)-1]) // HL

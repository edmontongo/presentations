package main

import "fmt"

// START OMIT
func TCReverse(acc, str string) string {
	if len(str) == 0 {
		return acc
	}
	return TCReverse(str[:1]+acc, str[1:]) // HL
}

func main() {
	const str = "0123456789"
	fmt.Println(str)
	fmt.Println(TCReverse("", str))
}

//END OMIT

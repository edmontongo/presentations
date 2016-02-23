package main

import "fmt"

func main() {
	a := 5
	b := &a
	fmt.Printf("a, &a: %x, %x\n", a, &a)
	fmt.Printf("b, &b, *b: %x, %x, %d\n", b, &b, *b)
	*b = 42
	fmt.Printf("a, &a: %x, %x\n", a, &a)
	fmt.Printf("b, &b, *b: %x, %x, %d\n", b, &b, *b)
}

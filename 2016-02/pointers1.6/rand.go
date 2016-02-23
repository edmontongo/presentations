package main

// #include <stdlib.h>
import "C"
import "fmt"

func main() {
	C.srandom(C.uint(42))
	i := int(C.random())
	fmt.Println(i)
}

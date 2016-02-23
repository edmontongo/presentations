package main

import (
	"fmt"
	"unsafe"
)

// #include <stdio.h>
// void printAddress(void* v) {
// 	printf("%x\n", v);
// }
import "C"

func main() {
	good := 42
	fmt.Printf("%p\n", &good)
	C.printAddress(unsafe.Pointer(&good))

	bad := struct{ i *int }{ &good }
	fmt.Printf("%p\n", &bad)
	defer func() { if r := recover(); r != nil { fmt.Println("Recovered in f", r) } }()
	C.printAddress(unsafe.Pointer(&bad))
}

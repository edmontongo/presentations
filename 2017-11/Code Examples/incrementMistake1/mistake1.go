package main

import "fmt"

func addToInt(A *int) {
	*A++
}

func main() {
	A := 0
	go func() {
		for i := 0; i < 10000000; i++ {
			addToInt(&A)
		}
	}()
	fmt.Println(A)
}

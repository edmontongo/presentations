package main

import "fmt"

func greet(who string) {
	fmt.Println("Hello,", who)
}

func main() {
	go greet("Go") // HL
}

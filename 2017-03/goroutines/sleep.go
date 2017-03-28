package main

import (
	"fmt"
	"time"
)

func greet(who string) {
	fmt.Println("Hello,", who)
}

func main() {
	go greet("sleep")

	time.Sleep(time.Second) // HL
}

package main

import (
	"fmt"
	"time"
)

func play(s string, n int) {
	for i := 1; i <= n; i++ {
		fmt.Printf("%s (#%d)\n", s, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go play("ping", 5)
	go play("\tpong", 5)
	time.Sleep(10 * time.Second)
}

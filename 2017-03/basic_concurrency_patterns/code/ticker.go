package main

import (
	"fmt"
	"time"
)

func main() {
	var sec = time.NewTicker(time.Second)
	var dot = time.NewTicker(100 * time.Millisecond)
	var timeout = time.After(10 * time.Second)

	for {
		select {
		case <-sec.C:
			fmt.Printf("%s\n", time.Now())
		case <-dot.C:
			fmt.Printf(".")
		case <-timeout:
			fmt.Println("\nDone - timeout")
			return
		}
	}
}

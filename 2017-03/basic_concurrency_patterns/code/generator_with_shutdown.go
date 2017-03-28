package main

import (
	"context"
	"fmt"
	"time"
)

func genNumbers(ctx context.Context) <-chan uint { // HL
	var ch = make(chan uint)

	go func() { // HL
		var n uint
		defer close(ch) // HL
		for {
			select { // HL
			case ch <- n: // HL
				time.Sleep(50 * time.Millisecond) // OMIT
				n++
			case <-ctx.Done(): // HL
				fmt.Println("\nshutting down generator:", ctx.Err())
				return
			} // HL
		}
	}()
	return ch // HL
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 4*time.Second) // HL

	for num := range genNumbers(ctx) { // HL
		fmt.Printf("\x0c%d", num)
	} // HL
}

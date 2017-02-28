package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context, ctx
	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second))

	fmt.Println("Starting...")

	// The channel returned will be closed after the deadline
	<-ctx.Done()

	fmt.Println("Done!")
}

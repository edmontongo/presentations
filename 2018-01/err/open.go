package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("/etc/passwd")
	if err != nil {
		// handle error
		fmt.Fprintf(os.Stderr, "File opening failed: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	// ...use the file...
}

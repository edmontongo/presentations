package main

import (
	"errors"
	"fmt"
)

func foo() error { return errors.New("foo crashed") }

func baz() error {
	if err := foo(); err != nil { // HL
		return fmt.Errorf("baz issue: %s", err) // HL
	}
	return nil
}

func bar() error {
	if err := baz(); err != nil { // HL
		return fmt.Errorf("bar failed: %s", err) // HL
	}
	return nil
}

func main() {
	fmt.Println("ERROR:", bar())
}

package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
)

func main() {
	var input = bytes.NewBufferString("")
	var data = make([]byte, 1)
	_ = errors.New(fmt.Sprintf("unused, to pacify playground: %s", io.EOF))
	log.SetFlags(0)

	io.EOF = errors.New("EOF = Errors Occur Frequently") // HL

	if _, err := input.Read(data); err != nil {
		log.Fatalf("failed to read: %s", err)
	}
}

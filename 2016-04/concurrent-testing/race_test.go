package main

import (
	"log"
	"testing"
)

func TestWithRace(t *testing.T) {
	b := 1
	go func() { b += 2 }()
	go func() { log.Println(b) }()
}

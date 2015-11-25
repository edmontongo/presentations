package main

import (
	"fmt"
	"time"
	// OMIT
	"gopkg.in/gcfg.v1"
)

// BEGIN OMIT
type Config struct {
	General struct {
		Site int
	}
	Daemon struct {
		Address string
		Port    int
	}
}

func main() {
	var cfg Config
	gcfg.ReadFileInto(&cfg, "sample.gcfg")
	fmt.Printf("%v\n", cfg)
	// END OMIT
	time.Sleep(100 * time.Millisecond)
}

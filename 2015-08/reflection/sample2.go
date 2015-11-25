package main

import (
	"fmt"
	// OMIT
	"time"

	"gopkg.in/gcfg.v1"
)

// BEGIN OMIT
type (
	Config struct {
		General General
		Daemon  Daemon
	}
	General struct {
		Site int
	}
	Daemon struct {
		Address string
		Port    int
	}
)

func NewConfig() *Config {
	return &Config{General{Site: 1234}, Daemon{Address: "0.0.0.0", Port: 5678}}
}

// E2 OMIT
func main() {
	cfg := NewConfig()
	gcfg.ReadFileInto(cfg, "sample.gcfg")
	// END OMIT

	fmt.Printf("%v\n", cfg)
	time.Sleep(100 * time.Millisecond)
}

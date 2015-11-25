package main

import (
	"fmt"
	// OMIT
	"gopkg.in/gcfg.v1"
)

func setDefaults(s interface{}) error {
	return nil
}

// BEGIN OMIT
type Config struct {
	General struct {
		Site int `default:"1234"`
	}
	Daemon struct {
		Address string `default:"0.0.0.0"`
		Port    int    `default:"5678"`
	}
}

func main() {
	var cfg Config
	setDefaults(&cfg) // HL
	gcfg.ReadFileInto(&cfg, "sample.gcfg")
	// END OMIT
	fmt.Printf("%v\n", cfg)
}

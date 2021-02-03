package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

func main() {
	var envs struct {
		Task     string        `envconfig:"SOME_TASK" default:"My task"`
		Duration time.Duration `default:"20s"`
	}
	envconfig.Usage("MYAPP", &envs)
	// fmt.Printf("Run '%s' for %v", envs.Task, envs.Duration)
}

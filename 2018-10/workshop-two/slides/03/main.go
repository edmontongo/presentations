package main

import (
	"flag"
	"fmt"
	"os"
)

// start OMIT
func main() {
	var secret string
	flag.StringVar(&secret, "secret", "", "DarkSky API key (required)")
	port := flag.String("port", "3000", "Port service will listen on")
	_ = port
	flag.Parse()

	if len(secret) == 0 {
		secret = os.Getenv("secret")
	}
	if len(secret) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}
	fmt.Printf("API secret key:%s\n", secret)
}

// end OMIT

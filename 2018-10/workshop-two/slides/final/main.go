package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/edmontongo/darksky"
)

func main() {
	var secret string
	// define and parse out command line flags
	flag.StringVar(&secret, "secret", "", "DarkSky API key (required)")
	port := flag.String("port", "3000", "Port the service will be listening on")
	flag.Parse()

	if len(secret) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// init
	tmplCache := newTemplateCache()
	forecastSvc := darksky.New(secret)
	locationHandler := newLocationHandler(forecastSvc, tmplCache)

	// handlers
	http.Handle("/", locationHandler)

	// start http server
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}

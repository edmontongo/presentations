package main

// svcstart OMIT
import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/edmontongo/darksky"
)

var darkskySvc *darksky.DarkSky

type locationHandler struct{}

func (lh locationHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// svcend OMIT

func main() {
	var secret string
	flag.StringVar(&secret, "secret", "", "DarkSky API key (required)")
	port := flag.String("port", "3000", "Port service will listen on")
	flag.Parse()

	if len(secret) == 0 {
		secret = os.Getenv("secret")
	}
	if len(secret) == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	// starthandle OMIT
	darkskySvc = darksky.New(secret)
	http.Handle("/", locationHandler{})

	fmt.Printf("Listening on :%s\n", *port)
	if err := http.ListenAndServe(":"+*port, nil); err != nil {
		fmt.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
	// endhandle OMIT
}

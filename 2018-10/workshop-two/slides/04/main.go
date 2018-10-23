package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// handle request
	})
	http.Handle("/foo", someHandler{})
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}

package main

import (
	"fmt"
	"log"
	"net/http"
)

var counter int

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		counter++
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, counter)
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		counter = 0
	})

	server := &http.Server{Addr: ":8080"}
	server.SetKeepAlivesEnabled(false)
	log.Fatal(server.ListenAndServe())
}

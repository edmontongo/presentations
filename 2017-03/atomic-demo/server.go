package main

import (
	"fmt"
	"log"
	"net/http"
	"sync/atomic"
)

var counter int64

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt64(&counter, 1)
	})

	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, atomic.LoadInt64(&counter))
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		atomic.StoreInt64(&counter, 0)
	})

	server := &http.Server{Addr: ":8080"}
	server.SetKeepAlivesEnabled(false)
	log.Fatal(server.ListenAndServe())
}

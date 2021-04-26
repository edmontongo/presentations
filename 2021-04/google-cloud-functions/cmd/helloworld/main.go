package main

import (
	"fmt"
	"net/http"

	"github.com/edmontongo/presentations/2021-04/google_cloud_functions"
)

var HelloWorld = google_cloud_functions.HelloWorld

func main() {
	fmt.Println("Serving on :8818...")
	http.HandleFunc("/", HelloWorld) // HL
	http.ListenAndServe(":8818", nil)
}

package main

import (
	"fmt"
	"net/http"

	"github.com/edmontongo/presentations/2021-04/google_cloud_functions"
)

var ExperimentParallelRequests = google_cloud_functions.ExperimentParallelRequests

func main() {
	fmt.Println("Serving on :8818...")
	http.HandleFunc("/", ExperimentParallelRequests) // HL
	http.ListenAndServe(":8818", nil)
}

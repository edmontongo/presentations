package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

// from https://github.com/edmontongo/workshopone 

func main() {
	get()
}

func get() {
	for {
		resp, err := http.Get("http://localhost:8080/")
		exitOnBadResponse(resp, err)
		b, err := ioutil.ReadAll(resp.Body)
		exitOnError(err)
		fmt.Println(string(b))
	}
}

func exitOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func exitOnBadResponse(resp *http.Response, err error) {
	exitOnError(err)
	if resp.StatusCode != http.StatusOK {
		panic(resp.Status)
	}
}

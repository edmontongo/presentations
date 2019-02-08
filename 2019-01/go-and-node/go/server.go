package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func fetch(url string) (string, error) {
	r, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

const url = "https://www.google.com/humans.txt"

func handler(w http.ResponseWriter, req *http.Request) {
	data, err := fetch(url)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusBadGateway)
		io.WriteString(w, "Bad gateway")
		return
	}
	io.WriteString(w, data)
}

func main() {
	http.HandleFunc("/", handler)

	const addr = ":3000"
	log.Printf("Listening on %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

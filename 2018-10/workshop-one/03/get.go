package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Lat, Long float64
}

func Forecast(host, secretKey string, l Location) {
	url := fmt.Sprintf("%v/forecast/%v/%v,%v", host, secretKey, l.Lat, l.Long)

	r, err := http.Get(url)
	if err != nil {
		// handle error
	}
	defer r.Body.Close()

	fmt.Println(r.StatusCode)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(b))
}

func main() {
	const host = "https://api.darksky.net"
	const secretKey = "0123456789abcdef9876543210fedcba"

	edmonton := Location{Lat: 53.5458874, Long: -113.5034304}
	Forecast(host, secretKey, edmonton)
}

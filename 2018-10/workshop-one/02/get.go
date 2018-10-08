package main

// BEGINIMPORT OMIT
import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// ENDIMPORT OMIT

type Location struct {
	Lat, Long float64
}

// BEGINFORECAST OMIT
func Forecast(secretKey string, l Location) {
	url := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v", secretKey, l.Lat, l.Long)

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

// ENDFORECAST OMIT

func main() {
	const secretKey = "0123456789abcdef9876543210fedcba"

	edmonton := Location{Lat: 53.5458874, Long: -113.5034304}
	Forecast(secretKey, edmonton)
}

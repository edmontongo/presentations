package main

import "fmt"

// BEGIN OMIT
type Location struct {
	Lat, Long float64
}

// END OMIT

// BEGINFORECAST OMIT
func Forecast(secretKey string, l Location) {
	// ENDFORECAST OMIT
	url := fmt.Sprintf("https://api.darksky.net/forecast/%v/%v,%v", secretKey, l.Lat, l.Long)
	fmt.Println(url)
}

func main() {
	const secretKey = "0123456789abcdef9876543210fedcba"

	edmonton := Location{Lat: 53.5458874, Long: -113.5034304}
	Forecast(secretKey, edmonton)
}

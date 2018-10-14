package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location struct {
	Lat, Long float64
}

type DarkSky struct {
	Host, SecretKey string
}

type Weather struct {
	Current Current `json:"currently"`
}

type Current struct {
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
}

var (
	ErrUnauthorized = errors.New("permission denied")
	ErrUnknown      = errors.New("unknown")
)

func NewDarkSky(host, secretKey string) *DarkSky {
	return &DarkSky{Host: host, SecretKey: secretKey}
}

func (ds *DarkSky) Forecast(l Location) (*Weather, error) {
	url := fmt.Sprintf("%v/forecast/%v/%v,%v", ds.Host, ds.SecretKey, l.Lat, l.Long)

	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	if r.StatusCode != http.StatusOK {
		switch r.StatusCode {
		case http.StatusUnauthorized:
			return nil, ErrUnauthorized
		default:
			return nil, ErrUnknown
		}
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	w := &Weather{}
	err = json.Unmarshal(b, &w)
	return w, err
}

func main() {
	const host = "https://api.darksky.net"
	const secretKey = "0123456789abcdef9876543210fedcba"
	ds := NewDarkSky(host, secretKey)

	edmonton := Location{Lat: 53.5458874, Long: -113.5034304}
	w, err := ds.Forecast(edmonton)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", w)
}

package main

import (
	"encoding/json"
	"log"
)

type Weather struct {
	Currently Currently `json:"currently"`
}

type Currently struct {
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
}

var data = []byte(`{
  "latitude": 53.5458874,
  "longitude": -113.5034304,
  "timezone": "America/Edmonton",
  "currently": {
    "summary": "Light Snow",
    "temperature": 20.33
  }
}`)

func main() {
	w := &Weather{}
	err := json.Unmarshal(data, &w)
	if err != nil {
		log.Fatal(err)
	}

	summary := w.Currently.Summary
	temperature := w.Currently.Temperature
	log.Printf("%v and %vºF", summary, temperature)
}

package main

import (
	"encoding/json"
	"fmt"
)

type Weather struct {
	Currently Currently `json:"currently"`
}

type Currently struct {
	Summary     string  `json:"summary"`
	Temperature float64 `json:"temperature"`
}

func main() {
	data := []byte(`{"currently":{"summary":"Mostly Cloudy","temperature":32.94}}`)
	w := Weather{}
	_ = json.Unmarshal(data, &w)
	fmt.Printf("%+v\n", w)
}

package main

import (
	"fmt"
)

func main() {
	var c, f float64
	c = 0.0
	f = CelsiusToFahrenheit(c)
	fmt.Printf("%vºC is %vºF\n", c, f)
}

// CelsiusToFahrenheit converts temperature for our American friends.
func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

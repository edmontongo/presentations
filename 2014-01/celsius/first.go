package main

import (
	"fmt"
)

func main() {
	var c, f float64
	c = -1.0
	f = CelsiusToFahrenheit(c)
	fmt.Printf("%4.2fºC is %4.2fºF\n", c, f)
}

// CelsiusToFahrenheit converts temperature to share with our American friends.
func CelsiusToFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

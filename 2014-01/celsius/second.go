package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main() {
	c := Celsius(-1.0)
	f := CelsiusToFahrenheit(c)
	fmt.Printf("%.2fºC is %.2fºF\n", c, f)

	// if c == f {
	// 	fmt.Println("Match")
	// }
}

// CelsiusToFahrenheit converts temperature for our American friends.
func CelsiusToFahrenheit(c Celsius) Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

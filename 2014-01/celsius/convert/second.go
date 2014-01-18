package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main() {
	c := Celsius(-1.0)
	f := CelsiusToFahrenheit(c)
	fmt.Printf("%4.2fºC is %4.2fºF\n", c, f)

	// if c == f {
	// 	fmt.Println("Match")
	// }
}

// CelsiusToFahrenheit converts temperature to share with our American friends.
func CelsiusToFahrenheit(t Celsius) Fahrenheit {
	return Fahrenheit((t * 9 / 5) + 32)
}

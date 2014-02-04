package main

import (
	"fmt"
)

type Celsius float64
type Fahrenheit float64

func main() {
	c := Celsius(-1.0)
	fmt.Printf("%.2fºC is %.2fºF\n", c, c.ToFahrenheit())
}

// ToFahrenheit converts temperature for our American friends.
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

// ToCelsius converts temperature from our American friends.
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

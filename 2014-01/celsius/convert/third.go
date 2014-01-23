package main

import (
	"fmt"
)

type Celsius float32
type Fahrenheit float32

func main() {
	c := Celsius(-1.0)
	fmt.Printf("%4.2fºC is %4.2fºF\n", c, c.ToFahrenheit())
}

// ToFahrenheit converts temperature to share with our American friends.
func (c Celsius) ToFahrenheit() Fahrenheit {
	return Fahrenheit((c * 9 / 5) + 32)
}

// ToCelsius converts temperature from our American friends.
func (f Fahrenheit) ToCelsius() Celsius {
	return Celsius((f - 32) * 5 / 9)
}

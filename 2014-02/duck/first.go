package main

import "fmt"

type Duck struct{}

func (duck Duck) Speak() string {
	return "Quack!"
}

func main() {
	donald := Duck{}
	fmt.Println(donald.Speak())
}

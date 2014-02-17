package main

import "fmt"

type Duck struct{}

func (duck Duck) Talk() string {
	return "Quack!"
}

// If it quacks like a duck...
func IsDuck(duck Duck) bool {
	return duck.Talk() == "Quack!"
}

func main() {
	donald := Duck{}
	fmt.Println(donald.Talk())
	fmt.Println("Is a duck?", IsDuck(donald))
}

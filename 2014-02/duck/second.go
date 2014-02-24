package main

import "fmt"

type Duck struct{}

func (duck Duck) Speak() string {
	return "Quack!"
}

type Dog struct{}

func (dog Dog) Speak() string {
	return "\U0001F436 LOL"
}

// If it quacks like a duck...
func IsDuck(s Speaker) bool {
	return s.Speak() == "Quack!"
}

type Speaker interface {
	Speak() string
}

func main() {
	donald := Duck{}
	fmt.Println(donald.Speak())
	fmt.Println("Is a duck?", IsDuck(donald))

	laughingDog := Dog{}
	fmt.Println(laughingDog.Speak())
	fmt.Println("Is a duck?", IsDuck(laughingDog))
}

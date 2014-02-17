package main

import "fmt"

type Duck struct{}

func (duck Duck) Talk() string {
	return "Quack!"
}

type Dog struct{}

func (dog Dog) Talk() string {
	return "🐶 LOL"
}

// If it quacks like a duck...
func IsDuck(s Talker) bool {
	return s.Talk() == "Quack!"
}

type Talker interface {
	Talk() string
}

type ClayPigeon struct{}

func main() {
	donald := Duck{}
	fmt.Println(donald.Talk())
	fmt.Println("Is a duck?", IsDuck(donald))

	laughingDog := Dog{}
	fmt.Println(laughingDog.Talk())
	fmt.Println("Is a duck?", IsDuck(laughingDog))

	clay := ClayPigeon{}
	fmt.Println("Is a duck?", IsDuck(clay))
}

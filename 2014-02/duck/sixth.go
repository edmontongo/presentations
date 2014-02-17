package main

import "fmt"

type Duck struct{}

func (duck Duck) Talk() string {
	return "Quack!"
}

func (duck Duck) Fly() {
	fmt.Println("flap flap")
}

type Dog struct{}

func (dog Dog) Talk() string {
	return "🐶 LOL"
}

// If it is a duck...
func IsDuck(s Talker) bool {
	duck, ok := s.(Duck)
	duck.Fly() // don't use type assertions for evil
	return ok
}

type Talker interface {
	Talk() string
}

func main() {
	donald := Duck{}
	fmt.Println(donald.Talk())
	fmt.Println("Is a duck?", IsDuck(donald))

	laughingDog := Dog{}
	fmt.Println(laughingDog.Talk())
	fmt.Println("Is a duck?", IsDuck(laughingDog))
}

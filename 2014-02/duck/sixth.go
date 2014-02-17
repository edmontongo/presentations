package main

import "fmt"

type Duck struct{}

func (duck Duck) Speak() string {
	return "Quack!"
}

func (duck Duck) Fly() string {
	return "flap flap"
}

type Dog struct{}

func (dog Dog) Speak() string {
	return "🐶 LOL"
}

func IsDuck(s Speaker) bool {
	switch t := s.(type) {
	default:
		return false
	case Duck:
		fmt.Println(t.Fly())
		return true
	}
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

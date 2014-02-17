package main

import "fmt"

type Duck struct{}

func (duck Duck) Talk() string {
	return "Quack!"
}

func (duck Duck) Fly() string {
	return "flap flap"
}

type Dog struct{}

func (dog Dog) Talk() string {
	return "🐶 LOL"
}

// If it is a duck...
func IsDuck(s Talker) bool {
	switch t := s.(type) {
	default:
		return false
	case Duck:
		fmt.Println(t.Fly())
		return true
	}
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

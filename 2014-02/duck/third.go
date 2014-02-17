package main

import "fmt"

type Duck struct{}

func (duck Duck) Speak() string {
	return "Quack!"
}

type Dog struct{}

func (dog Dog) Speak() string {
	return "🐶 LOL"
}

func IsDuck(duck Duck) bool {
	return duck.Speak() == "Quack!"
}

func main() {
	donald := Duck{}
	fmt.Println(donald.Speak())
	fmt.Println("Is a duck?", IsDuck(donald))

	laughingDog := Dog{}
	fmt.Println(laughingDog.Speak())
}

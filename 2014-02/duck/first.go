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

func main() {
	donald := Duck{}
	fmt.Println(donald.Speak())

	laughingDog := Dog{}
	fmt.Println(laughingDog.Speak())
}

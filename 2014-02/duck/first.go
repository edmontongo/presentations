package main

import "fmt"

type Duck struct{}

func (duck Duck) Talk() string {
	return "Quack!"
}

func main() {
	donald := Duck{}
	fmt.Println(donald.Talk())
}

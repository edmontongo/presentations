package main

import (
	"fmt"
)

type Stack struct {
	top    *Element
	length int
}

type Element struct {
	value interface{}
	next  *Element
}

// Len returns the length of the stack.
func (s *Stack) Len() int {
	return s.length
}

// Push a new element onto the stack.
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value: value, next: s.top}
	s.length++
}

// Pop the top element from the stack and return its value.
// If the stack is empty, return nil.
func (s *Stack) Pop() (value interface{}) {
	if s.length > 0 {
		value, s.top = s.top.value, s.top.next
		s.length--
		return value
	}
	return nil
}

type Duck struct{}

func (duck Duck) Speak() string {
	return "Quack!"
}

func main() {
	stack := new(Stack)

	stack.Push("LOL")
	stack.Push(Duck{})

	for stack.Len() > 0 {
		switch s := stack.Pop().(type) {
		case Duck:
			fmt.Println(s.Speak())
		case string:
			fmt.Println(s)
		}
	}
}

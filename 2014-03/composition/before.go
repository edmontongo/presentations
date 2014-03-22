package main

import "fmt"

type Person struct {
	FirstName, LastName string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

type Employee struct {
	person Person
}

func main() {
	employee := Employee{
		Person{FirstName: "Daniel", LastName: "Huckstep"},
	}
	fmt.Println(employee.person.Name())
}

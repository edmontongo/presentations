package main

import "fmt"

type Person struct {
	FirstName, LastName string
}

func (p *Person) Name() string {
	return p.FirstName + " " + p.LastName
}

type Employee struct {
	Person
	Department
}

type Department struct {
	Name string
}

func (e *Employee) Name() string {
	return e.Person.Name() + " / " + e.Department.Name
}

func main() {
	employee := Employee{
		Person{FirstName: "Daniel", LastName: "Huckstep"},
		Department{Name: "IT"},
	}
	fmt.Println(employee.Name())
}

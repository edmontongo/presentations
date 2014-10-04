package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	// Basics
	fmt.Printf("Press any enter key to start")
	wait()
	cursorSimple()
	cursorStruct()
	cursorMove()
	conditional()
	methodSimple()
}

func execute(templ string, data interface{}) {
	fmt.Printf("Template:\t%s\nData:    \t%T: %#v\n", strings.Replace(templ, "\n", "\\n", -1), data, data)
	wait()
	err := template.Must(template.New("s").Parse(templ)).Execute(os.Stdout, data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println()
	fmt.Println()
	wait()
}

var buf = bufio.NewReader(os.Stdin)

func wait() {
	buf.ReadLine()
}

// The cursor refers to the current object
func cursorSimple() {
	fmt.Println("Anything inside of the '{{}}'' is called a pipeline, similar in spirit to a Unix pipeline. The '.' inside is called a cursor and refers to the current value")
	data := "Hello!"
	execute("{{.}}", data)
}

// The cursor automatically resolves your type into a string
func cursorStruct() {
	data := struct {
		Int    int
		Float  float32
		String string
	}{9, 3.14159, "Purple"}

	fmt.Println("Here we refer to members of '.'.")
	execute("{{.Int}} bears ate {{.Float}} {{.String}} coconuts", data)
}

// The cursor can move
func cursorMove() {
	fmt.Println("Range acts on a map or slice, and executes the template until end on every member of that data structure with the cursor reset as one would expect.")
	execute("{{.}}:\n{{range .}}\t{{.}} bottles of beer on the wall, {{.}} bottles of beer!\n{{end}}", []int{99, 98, 97})
}

// Conditionals are similar to duck typed booleans
func conditional() {
	i := 5
	data := []interface{}{
		true, false,
		0, 1,
		"", "Hi",
		&i, nil,
		[]string{"one", "two"}, []string{}}

	fmt.Println("The if conditional uses fuzzy false. '', 0, nil, false all evaluate as false.")
	execute("{{range .}}{{if .}}True!{{else}}False{{end}} -- {{.}}\n{{end}}", data)
}

// You can call methods, pass them arguments, pass their results into other methods, etc.
func methodSimple() {
	fmt.Println("This complicated pipeline passes the result of the first function as an argument to the second, resulting in a Cube operator.")
	execute("{{ .Times . | .Times}}", Number(64))
}

type Number int

func (n Number) Times(m Number) Number {
	return n * m
}

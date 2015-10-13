package main

import (
	"fmt"
	"go/scanner"
	"go/token"
)

func main() {
	scanner := scanner.Scanner{}

	file, data := stringAsFile("position := initial + rate * 60\n")
	scanner.Init(file, data, nil, 0)
	for pos, tok, lit := scanner.Scan(); tok != token.EOF; pos, tok, lit = scanner.Scan() {
		fmt.Printf("token: %v, %v, '%v'\n", pos, tok, lit)
	}
}

func stringAsFile(s string) (*token.File, []byte) {
	fs := token.NewFileSet()
	input := []byte(s)
	file := fs.AddFile("lexer-sample.go", fs.Base(), len(input))
	return file, input
}

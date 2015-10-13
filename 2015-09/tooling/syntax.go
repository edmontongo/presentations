package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "tooling/syntax-sample.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// BEGIN OMIT
	// Print the names and locations of all top level functions
	for _, decl := range file.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			pos := fset.Position(funcDecl.Pos())
			fmt.Printf("%s:% 3d: %s\n", pos.Filename, pos.Line, funcDecl.Name)
		}
	}
	// END OMIT
}

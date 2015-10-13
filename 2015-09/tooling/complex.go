package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type funcVisitor struct {
	*token.FileSet
	*ast.FuncDecl
}

func (fv *funcVisitor) Visit(node ast.Node) ast.Visitor {
	if stmt, ok := node.(*ast.IfStmt); ok {
		if r := fv.Recv; r != nil {
			fmt.Printf("(%v) ", r.List[0].Type.(*ast.StarExpr).X)
		}
		funcLine := fv.Position(fv.Pos()).Line
		ifLine := fv.Position(node.Pos()).Line
		fmt.Printf("%s:%d-%d:\t", fv.Name, funcLine, ifLine-funcLine)

		fmt.Printf("If statement ")
		if stmt.Init != nil {
			fmt.Printf("with initialization ")
		}
		fmt.Println()
	}
	return fv
}

// ENDFUNC OMIT

type basicVisitor struct{ *token.FileSet }

func (bv *basicVisitor) Visit(node ast.Node) ast.Visitor {

	// Only process functions
	if funcDecl, ok := node.(*ast.FuncDecl); ok {
		return &funcVisitor{bv.FileSet, funcDecl}
	}

	return nil
}

// ENDBASIC OMIT

func main() {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, "tooling/complex.go", nil, 0)
	if err != nil {
		fmt.Println(err)
		return
	}

	// BEGIN OMIT
	// Run a basicVisitor on all
	for _, d := range file.Decls {
		ast.Walk(&basicVisitor{fileSet}, d)
	}
	// END OMIT
}

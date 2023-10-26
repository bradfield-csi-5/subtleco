package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

// Given an AST node corresponding to a function (guaranteed to be
// of the form `func f(x, y byte) byte`), compile it into assembly
// code.
//
// Recall from the README that the input parameters `x` and `y` should
// be read from memory addresses `1` and `2`, and the return value
// should be written to memory address `0`.
func compile(node *ast.FuncDecl) (string, error) {
	var assembly string
	body := node.Body.List
	for _, stmt := range body {
		switch x := stmt.(type) {
		case *ast.ReturnStmt:
			for _, result := range x.Results {
				switch y := result.(type) {
				case *ast.BasicLit:
					fmt.Println(y.Value)
					assembly += "pushi " + y.Value + "\npop 0\n"
				case *ast.BinaryExpr:
					assembly += "push 2\npush 1\n"
					switch y.Op {
					case token.ADD:
						assembly += "add\n"
					case token.SUB:
						assembly += "sub\n"
					case token.MUL:
						assembly += "mul\n"
					case token.QUO:
						assembly += "div\n"
					}
					assembly += "pop 0\n"
					fmt.Printf("%T\n", y)
				}
			}
		}
	}
	// fset := token.NewFileSet()
	// ast.Print(fset, body)
	assembly += "halt\n"
	return assembly, nil
}

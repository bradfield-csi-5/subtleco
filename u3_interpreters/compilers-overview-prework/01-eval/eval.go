package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"strconv"
)

// Given an expression containing only int types, evaluate
// the expression and return the result.
func Evaluate(expr ast.Expr) (int, error) {
	switch xT := expr.(type) {

	case *ast.BasicLit:
		return strconv.Atoi(xT.Value)

	case *ast.BinaryExpr:
		x, _ := Evaluate(xT.X)
		y, _ := Evaluate(xT.Y)
		op := xT.Op
		return arith(x, y, op)

	case *ast.ParenExpr:
		x, _ := Evaluate(xT.X)
		return x, nil

	default:
		fset := token.NewFileSet()
		_ = ast.Print(fset, xT)
	}
	return 0, nil
}

func arith(x, y int, op token.Token) (int, error) {
	switch op {
	case token.ADD:
		return x + y, nil
	case token.SUB:
		return x - y, nil
	case token.MUL:
		return x * y, nil
	default:
		fmt.Println("op is ", op)
		return 0, nil
	}
}

func main() {
	expr, err := parser.ParseExpr("1 + 2 - 3 * 4")
	if err != nil {
		log.Fatal(err)
	}
	fset := token.NewFileSet()
	err = ast.Print(fset, expr)
	if err != nil {
		log.Fatal(err)
	}
}

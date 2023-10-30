package main

import (
	"fmt"
	"log"
	"os"
)

type ExprType string

const (
	E_AND  ExprType = "AND"
	E_OR            = "OR"
	E_NOT           = "NOT"
	E_TERM          = "TERM"
)

type Expr interface {
	String() string
}

type Empty struct{}

func (e *Empty) String() string {
	return ""
}

type Term struct {
	Identifier string
}

func (t *Term) String() string {
	return fmt.Sprintf("TERM(%s)", t.Identifier)
}

type And struct {
	Left  Expr
	Right Expr
}

func (a *And) String() string {
	return fmt.Sprintf("AND(%s,%s)", a.Left.String(), a.Right.String())
}

type Or struct {
	Left  Expr
	Right Expr
}

func (o *Or) String() string {
	return fmt.Sprintf("OR(%s,%s)", o.Left.String(), o.Right.String())
}

type Not struct {
	Right Expr
}

func (n *Not) String() string {
	return fmt.Sprintf("NOT(%s)", n.Right.String())
}

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Why have you done this stop putting so many things in")
	}
	input := os.Args[1]
	var terms []Expr
	var ops []Expr
	tokens := Scan(input)
	for _, token := range tokens {
		switch token.Type {
		case IDENTIFIER:
			if len(ops) > 0 {
				switch oT := ops[0].(type) {
				case *Or:
					oT.Right = &Term{Identifier: token.Lexeme}
				case *And:
					oT.Right = &Term{Identifier: token.Lexeme}
				case *Not:
					oT.Right = &Term{Identifier: token.Lexeme}
				default:
					fmt.Println("HOW ARE YOU HERE")
					os.Exit(1)
				}
			}
			if len(ops) == 0 {
				terms = append(terms, &Term{Identifier: token.Lexeme})
			}
		case AND:
			if len(ops) == 1 {
				ops[0] = &And{Left: ops[0]}
			}
		case OR:
			if len(ops) == 1 {
				ops[0] = &Or{Left: ops[0]}
			}
		case NOT:
		default:
			fmt.Println("HOW DID YOU GET HERE")
		}
	}
	fmt.Println(ops[0].String())
}

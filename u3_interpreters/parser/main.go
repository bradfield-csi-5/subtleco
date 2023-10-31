package main

import (
	"fmt"
	"os"
)

/* lox parser grammar

expression -> equality;
equality -> comparison ( ( "!=" | "==" ) comparison )*;
comparison -> term ( ( ">" | ">=" | "<" | "<=" ) term )*;
term -> factor ( ( "-" | "+" ) factor )*;
factor -> unary ( ( "/" | "*" ) unary )*;
unary -> ( "!" | "-") unary | primary;
primary -> NUMBER | STRING | "true" | "false" | "nil" | "(" expression ")";
*/

/* My parser grammar
expression -> or
or -> and ( "OR" and)*
and -> not ( "AND" not)*
not -> "NOT" not | term
term -> STRING
*/

type Parser struct {
	Tokens  []Token
	Current int
}

func (p *Parser) previous() Token {
	return p.Tokens[p.Current-1]
}

func (p *Parser) peek() Token {
	return p.Tokens[p.Current]
}

func (p *Parser) isAtEof() bool {
	return p.peek().Type == EOF
}

func (p *Parser) advance() Token {
	if !p.isAtEof() {
		p.Current++
	}
	return p.previous()
}

func (p *Parser) check(tType TokenType) bool {
	if p.isAtEof() {
		return false
	}
	return p.peek().Type == tType
}

func (p *Parser) match(tTypes ...TokenType) bool {
	for _, tType := range tTypes {
		if p.check(tType) {
			p.advance()
			return true
		}
	}
	return false
}

func (p *Parser) parse() Expr {
	return p.expression()
}

func (p *Parser) expression() Expr {
	return p.or()
}

func (p *Parser) or() Expr {
	expr := p.and()
	for p.match(OR) {
		right := p.and()
		expr = &Or{
			Left:  expr,
			Right: right,
		}
		return expr
	}

	return expr
}

func (p *Parser) and() Expr {
	expr := p.not()
	for p.match(AND) {
		right := p.not()
		expr = &And{
			Left:  expr,
			Right: right,
		}
		return expr
	}
	return expr
}

func (p *Parser) not() Expr {
	if p.match(NOT) {
		right := p.not()
		return &Not{Right: right}
	}

	return p.term()
}

func (p *Parser) term() Expr {
	p.advance()
	term := &Term{
		Identifier: p.previous().Lexeme,
	}
	return term
}

type Term struct {
	Identifier string
}

func (t *Term) String() string {
	return fmt.Sprintf("TERM(%s)", t.Identifier)
}

type Not struct {
	Right Expr
}

func (n *Not) String() string {
	return fmt.Sprintf("NOT(%s)", n.Right.String())
}

type And struct {
	Left  Expr
	Right Expr
}

func (a *And) String() string {
	return fmt.Sprintf("AND(%s, %s)", a.Left.String(), a.Right.String())
}

type Or struct {
	Left  Expr
	Right Expr
}

func (o *Or) String() string {
	return fmt.Sprintf("OR(%s, %s)", o.Left.String(), o.Right.String())
}

type Expr interface {
	String() string
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("NO WHY NO")
	}

	tokens := Scan(os.Args[1])
	parser := Parser{
		Tokens:  tokens,
		Current: 0,
	}
	ast := parser.parse()
	fmt.Println(ast)
}

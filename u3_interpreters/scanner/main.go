package main

import (
	"fmt"
	"log"
	"os"
)

type TokenType string

const (
	AND        TokenType = "AND"
	OR                   = "OR"
	NOT                  = "NOT"
	IDENTIFIER           = "IDENTIFIER"
)

var keywords = map[string]TokenType{
	"AND": AND,
	"OR":  OR,
	"NOT": NOT,
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Line    int
	Literal any
}

type Scanner struct {
	Input   string
	Start   int
	Current int
	Line    int
	Length  int
	Tokens  []Token
}

func (s *Scanner) scan() {
	for {
		if s.Current >= s.Length {
			break
		}
		c := s.advance()

		switch c {

		case " ":
			s.Start++
		case "\n":
			s.Line++

		default:
			if isAlpha(c) {
				s.identifier()
			}
		}
	}
}

func (s *Scanner) advance() string {
	char := string(s.Input[s.Current])
	s.Current++
	return char
}

func (s *Scanner) peek() string {
	if s.isAtEnd() {
		return "\\0"
	}
	return string(s.Input[s.Current])
}

func (s *Scanner) identifier() {
	for {
		if isAlphaNumeric(s.peek()) {
			s.advance()
		} else {
			text := s.Input[s.Start:s.Current]
			ttype, ok := keywords[text]
			if !ok {
				ttype = IDENTIFIER
			}
			s.addToken(ttype)
			break
		}
	}
}

func (s *Scanner) isAtEnd() bool {
	return s.Current >= s.Length
}

func (s *Scanner) addToken(ttype TokenType) {
	newToken := Token{
		Type:    ttype,
		Line:    s.Line,
		Lexeme:  s.Input[s.Start:s.Current],
		Literal: s.Input[s.Start:s.Current],
	}
	s.Tokens = append(s.Tokens, newToken)
	s.Start = s.Current
}

// hello AND world OR alice AND NOT bob
func isAlpha(c string) bool {
	return (c >= "a" && c <= "z") ||
		(c >= "A" && c <= "Z") ||
		c == "_"
}

func isDigit(c string) bool {
	return c >= "0" && c <= "9"
}

func isAlphaNumeric(c string) bool {
	return isAlpha(c) || isDigit(c)
}

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Why have you done this stop putting so many things in")
	}
	scanner := Scanner{
		Input:   os.Args[1],
		Start:   0,
		Current: 0,
		Line:    1,
		Length:  len(os.Args[1]),
	}

	scanner.scan()
	for _, token := range scanner.Tokens {
		fmt.Printf("%+v\n", token)
	}
}

package main

type TokenType string

const (
	AND        TokenType = "AND"
	OR                   = "OR"
	NOT                  = "NOT"
	IDENTIFIER           = "IDENTIFIER"
	EOF                  = "EOF"
)

var keywords = map[string]TokenType{
	"AND": AND,
	"OR":  OR,
	"NOT": NOT,
	"EOF": EOF,
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

func Scan(input string) []Token {
	scanner := Scanner{
		Input:   input,
		Start:   0,
		Current: 0,
		Line:    1,
		Length:  len(input),
	}

	scanner.scan()
	eof := Token{
		Type:    EOF,
		Line:    scanner.Line,
		Lexeme:  "EOF",
		Literal: "EOF",
	}
	scanner.Tokens = append(scanner.Tokens, eof)
	return scanner.Tokens
}

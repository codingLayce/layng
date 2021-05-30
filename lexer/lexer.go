package lexer

import "layng/token"

// TODO: Use io.Reader instead of string for input in order to take source code from file
// TODO: Support the use of full Unicode (rework characters reading)

type Lexer struct {
	input        string
	position     int // point to current char
	readPosition int // point after current char
	currentChar  byte
}

/* ---------- EXPORTED FUNCTIONS ---------- */

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.currentChar {
	case '=':
		tok = newToken(token.ASSIGN, l.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, l.currentChar)
	case '(':
		tok = newToken(token.LPAREN, l.currentChar)
	case ')':
		tok = newToken(token.RPAREN, l.currentChar)
	case ',':
		tok = newToken(token.COMMA, l.currentChar)
	case '+':
		tok = newToken(token.PLUS, l.currentChar)
	case '{':
		tok = newToken(token.LBRACE, l.currentChar)
	case '}':
		tok = newToken(token.RBRACE, l.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

/* ---------- HELPERS FUNCTIONS ---------- */

// Read the next character and advance the cursor to the next one
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // ASCII code for "NUL"
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

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

// Lexer constructor
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Reads the next token and returns it
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

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
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}
		if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		}
		tok = newToken(token.ILLEGAL, l.currentChar)
	}

	l.readChar()
	return tok
}

/* ---------- PRIVATE FUNCTIONS ---------- */

// Reads the whole number until it reaches a non digit character
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Reads the whole identifier until it reaches a non letter character
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// Reads the next character and advance the cursor to the next one
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.currentChar = 0 // ASCII code for "NUL"
	} else {
		l.currentChar = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

/* ---------- HELPER FUNCTIONS ---------- */

// Skips whitespaces, tabulation and new lines by ignoring them and advance the cursor
func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

// Checks whether or not the byte is a digit based on the ASCII code
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// Checks whether or not the byte is a letter based on the ASCII code
// !! '_' is considered as a letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= '2' || ch == '_'
}

// Creates a Token object with the given fields
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

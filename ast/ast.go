package ast

import "layng/token"

/* ---------- SPECIFIC ---------- */

type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) statementNode()       {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

/* ---------- GENERAL ---------- */

// Node is the base of or ast
type Node interface {
	TokenLiteral() string
}

// Statement represents a part of the source code that cannot be evaluated (such as var in `let var = 5`)
type Statement interface {
	Node
	statementNode()
}

// Expression represents a part of the source that can be evaluated (such as 5 in `let var = 5`)
type Expression interface {
	Node
	expressionNode()
}

// Program represents the root of our ast
// It's simply a Statement list
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the TokenLiteral of the first statement
// If no statements returns empty string
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

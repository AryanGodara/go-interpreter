package lox

import "fmt"

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"for":    FOR,
	"fun":    FUN,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Token interface {
	Token(tokenType TokenType, lexeme string, literal interface{}, line int) Token
	ToString() string
}

type token struct {
	tokenType TokenType
	lexeme    string
	literal   interface{}
	line      int
}

func (t token) Token(tokenType TokenType, lexeme string, literal interface{}, line int) Token {
	return &token{tokenType, lexeme, literal, line}
}

func (t token) ToString() string {
	return fmt.Sprintf("%v %s %s", t.tokenType, t.lexeme, t.literal)
}

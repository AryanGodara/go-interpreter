package lox

import (
	"fmt"
	"strconv"
)

type Scanner interface {
	ScanTokens() []Token
}

type scanner struct {
	source string
	tokens []Token

	start   int
	current int
	line    int
}

// Scanner returns a new scanner.
func New(source string) Scanner {
	return &scanner{source, []Token{}, 0, 0, 1}
}

// ScanTokens scans the source file and returns a slice of tokens.
func (s *scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		// We are at the beginning of the next lexeme.
		s.start = s.current
		s.scanToken()
	}

	s.tokens = append(s.tokens, token{EOF, "", nil, 1})
	return s.tokens
}

// isAtEnd returns true if the scanner has reached the end of the source file.
func (s *scanner) isAtEnd() bool {
	return s.current >= len(s.source)
}

// scanToken scans the next token in the source file.
func (s *scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PARAM)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(MINUS)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			// A comment goes until the end of the line.
			for s.peek() != '\n' && !s.isAtEnd() {
				_ = s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		// Ignore whitespace.
	case '\n':
		s.line++
	case '"':
		s.string()
	default:
		// The default case now works overtime to help detect alphanumerics :( [ Like and share this repo to make Sir default feel that their work is appreciated :D ]
		if s.isDigit(c) {
			s.number()
		} else if s.isAlpha(c) {
			s.identifier()
		} else {
			Error(s.line, "Unexpected character.")
		}
	}
}

// advance consumes the next character in the source file and returns it.
func (s *scanner) advance() byte {
	s.current++
	return s.source[s.current]
}

// match consumes the current character in the source file if it matches the expected character, returning true if it did, and advancing the current forward.
func (s *scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}
	s.current++
	return true
}

// peek returns the current character in the source file without consuming it.
func (s *scanner) peek() byte {
	if s.isAtEnd() {
		return '\000'
	}
	return s.source[s.current]
}

// peekNext returns the next character in the source file without consuming it.
func (s *scanner) peekNext() byte {
	if s.current+1 >= len(s.source) {
		return '\000'
	}
	return s.source[s.current+1]
}

// string consumes the next string in the source file and adds it to the tokens slice.
func (s *scanner) string() {
	for s.peek() != '"' && !s.isAtEnd() {
		if s.peek() == '\n' {
			s.line++
		}
		_ = s.advance()
	}

	// The Closing
	if s.isAtEnd() {
		Error(s.line, "Unterminated string.")
		return
	}

	// Trim the surrounding quotes
	value := s.source[s.start+1 : s.current-1]
	s.addTokenLiteral(STRING, value)
}

// number consumes the next number in the source file and adds it to the tokens slice.
func (s *scanner) number() {
	for s.isDigit(s.peek()) {
		_ = s.advance()
	}

	if s.peek() == '.' && s.isDigit(s.peekNext()) {
		// Consume the "."
		_ = s.advance()

		for s.isDigit(s.peek()) {
			_ = s.advance()
		}
	}

	s.addTokenLiteral(NUMBER, s.convertToFloat64(s.source[s.start:s.current]))
}

// identifier consumes the next identifier in the source file and adds it to the tokens slice.
func (s *scanner) identifier() {
	for s.isAlphaNumeric(s.peek()) {
		s.advance()
	}
	tokenType := keywords[s.source[s.start:s.current]]
	if tokenType == 0 { //! I'm not sure about why this is 0 and not nil, I think I should google it, but java is not my cup of tea, so I'll just leave it here
		tokenType = IDENTIFIER
	}
	s.addToken(tokenType)
}

// isDigit returns true if the given byte is a digit.
func (s *scanner) isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

// isAlpha returns true if the given byte is an alphabetic character.
func (s *scanner) isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

// isAlphaNumeric returns true if the given byte is an alphabetic or numeric character.
func (s *scanner) isAlphaNumeric(c byte) bool {
	return s.isAlpha(c) || s.isDigit(c)
}

// addToken adds a token to the tokens slice.
func (s *scanner) addToken(tokenType TokenType) {
	s.addTokenLiteral(tokenType, nil)
}

// addTokenLiteral adds a token to the tokens slice with a literal value.
func (s *scanner) addTokenLiteral(tokenType TokenType, literal interface{}) {
	text := s.source[s.start:s.current]
	s.tokens = append(s.tokens, token{tokenType, text, literal, s.line})
}

// function to convert string to float64
func (s *scanner) convertToFloat64(str string) float64 {

	// Convert string to float
	res, err := strconv.ParseFloat(str, 64)

	if err != nil {
		fmt.Println("string = ", str)
		Error(s.line, "Error converting string to float64")
	}

	return res
}

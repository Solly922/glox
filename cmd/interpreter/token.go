package main

import "fmt"

type TokenType int

const (
	// Single-character tokens.
	LEFT_PAREN TokenType = iota
	RIGHT_PAREN
	LEFT_BRACE
	RIGHT_BRACE
	COMMA
	DOT
	MINUS
	PLUS
	SEMICOLON
	SLASH
	STAR

	// One or two character tokens.
	BANG
	BANG_EQUAL
	EQUAL
	EQUAL_EQUAL
	GREATER
	GREATER_EQUAL
	LESS
	LESS_EQUAL

	// Literals.
	IDENTIFIER
	STRING
	NUMBER

	// Keywords.
	AND
	CLASS
	ELSE
	FALSE
	FUN
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE

	EOF
)

func (tt TokenType) String() string {
	return [...]string{
		"LEFT_PAREN", "RIGHT_PAREN", "LEFT_BRACE", "RIGHT_BRACE",
		"COMMA", "DOT", "MINUS", "PLUS", "SEMICOLON", "SLASH", "STAR",
		"BANG", "BANG_EQUAL", "EQUAL", "EQUAL_EQUAL",
		"GREATER", "GREATER_EQUAL", "LESS", "LESS_EQUAL",
		"IDENTIFIER", "STRING", "NUMBER",
		"AND", "CLASS", "ELSE", "FALSE", "FUN", "FOR", "IF", "NIL", "OR",
		"PRINT", "RETURN", "SUPER", "THIS", "TRUE", "VAR", "WHILE",
		"EOF",
	}[tt]
}

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
	Line      int
}

func NewToken(tokenType TokenType, lexeme string, literal interface{}, line int) *Token {
	return &Token{TokenType: tokenType, Lexeme: lexeme, Literal: literal, Line: line}
}

func (t *Token) String() string {
	return t.TokenType.String() + " " + t.Lexeme + " " + TokenLiteralToString(t.Literal) // this last section should be t.Literal (I don't understand how that one should be typed yet)
}

func TokenTypeToString(tokenType TokenType) string {
	switch tokenType {
	case LEFT_PAREN:
		return "("
	case RIGHT_PAREN:
		return ")"
	case LEFT_BRACE:
		return "{"
	case RIGHT_BRACE:
		return "}"
	case COMMA:
		return ","
	case DOT:
		return "."
	case MINUS:
		return "-"
	case PLUS:
		return "+"
	case SEMICOLON:
		return ";"
	case STAR:
		return "*"
	case EOF:
		return ""
	case BANG_EQUAL:
		return "!="
	case BANG:
		return "!"
	case EQUAL_EQUAL:
		return "=="
	case EQUAL:
		return "="
	case LESS_EQUAL:
		return "<="
	case LESS:
		return "<"
	case GREATER_EQUAL:
		return ">="
	case GREATER:
		return ">"
	case SLASH:
		return "/"

	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"

	// Keywords.
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUN:
		return "FUN"
	case FOR:
		return "FOR"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case WHILE:
		return "WHILE"

	default:
		return "N/A"
	}
}

func TokenLiteralToString(literal interface{}) string {
	switch literal := literal.(type) {
	case string:
		return literal
	case []rune:
		return string(literal)
	case nil:
		return "null" // this is to appease tests
	case float64:
		if literal == float64(int(literal)) {
			return fmt.Sprintf("%0.1f", literal) // this is to appease tests (ie 1234.0)
		} else {
			return fmt.Sprintf("%g", literal) // keeps precision for non-whole numbers
		}
	default:
		return fmt.Sprintf("%v", literal)
	}
}

package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING
	INDENTIFIER

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	ASSIGNMENT
	EQUALS
	NOT
	NOT_EQUALS

	LESS
	GREATER
	LESS_EQUALS
	GREATER_EQUALS
	OR
	AND

	DOT
	DOT_DOT // 0..10
	SEMI_COLON
	COLON
	QUESTION
	COMMA
	PERCENT

	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS

	PLUS
	MINUS
	DASH
	SLASH
	STAR

	// RESERVED KEYWORDS
	LET
	CONST
	CLASS
	NEW
	IMPORT
	FROM
	FN
	IF
	ELSE
	FOREACH
	WHILE
	FOR
	EXPORT
	TYPEOF
	IN
)

type Token struct {
	Value string
	Kind  TokenKind
}

func (token Token) isOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == token.Kind {
			return true
		}
	}
	return false
}

func (token Token) Debug() {
	if token.isOneOfMany(INDENTIFIER, NUMBER, STRING) {
		fmt.Printf("%s (%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s ()\n", TokenKindString(token.Kind))
	}
}

func NewToken(kind TokenKind, value string) Token {
	return Token{
		Kind:  kind,
		Value: value,
	}
}

// TokenKind is int and it returns string of what that int represents
func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "EOF"
	case NUMBER:
		return "NUMBER"
	case STRING:
		return "STRING"
	case INDENTIFIER:
		return "INDENTIFIER"
	case OPEN_BRACKET:
		return "OPEN_BRACKET"
	case CLOSE_BRACKET:
		return "CLOSE_BRACKET"
	case OPEN_CURLY:
		return "OPEN_CURLY"
	case CLOSE_CURLY:
		return "CLOSE_CURLY"
	case OPEN_PAREN:
		return "OPEN_PAREN"
	case CLOSE_PAREN:
		return "CLOSE_PAREN"
	case ASSIGNMENT:
		return "ASSIGNMENT"
	case EQUALS:
		return "EQUALS"
	case NOT_EQUALS:
		return "NOT_EQUALS"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	case LESS_EQUALS:
		return "LESS_EQUALS"
	case GREATER_EQUALS:
		return "GREATER_EQUALS"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case DOT:
		return "DOT"
	case DOT_DOT:
		return "DOT_DOT"
	case SEMI_COLON:
		return "SEMI_COLON"
	case COLON:
		return "COLON"
	case PERCENT:
		return "PERCENT"
	case QUESTION:
		return "QUESTION"
	case COMMA:
		return "COMMA"
	case PLUS_PLUS:
		return "PLUS_PLUS"
	case MINUS_MINUS:
		return "MINUS_MINUS"
	case PLUS_EQUALS:
		return "PLUS_EQUALS"
	case MINUS_EQUALS:
		return "MINUS_EQUALS"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case DASH:
		return "DASH"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case LET:
		return "LET"
	case CONST:
		return "CONST"
	case CLASS:
		return "CLASS"
	case NEW:
		return "NEW"
	case IMPORT:
		return "IMPORT"
	case FROM:
		return "FROM"
	case FN:
		return "FN"
	case IF:
		return "IF"
	case ELSE:
		return "ELSE"
	case FOREACH:
		return "FOREACH"
	case WHILE:
		return "WHILE"
	case EXPORT:
		return "EXPORT"
	case FOR:
		return "FOR"
	case TYPEOF:
		return "TYPEOF"
	case IN:
		return "IN"
	}
	return ""
}

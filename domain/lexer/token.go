package lexer

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER"
	NUMERAL    = "NUMERAL"

	ASSIGN = "="

	COMMA = ","

	OPEN_PARENTHESIS  = "("
	CLOSE_PARENTHESIS = ")"

	OPEN_BRACES  = "{"
	CLOSE_BRACES = "}"

	FUNCTION = "FUNCTION"
)

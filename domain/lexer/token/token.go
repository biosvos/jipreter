package token

type Kind string

type Token struct {
	kind    Kind
	literal string
}

func NewToken(kind Kind, literal string) *Token {
	return &Token{kind: kind, literal: literal}
}

const (
	ILLEGAL = Kind("ILLEGAL")
	EOF     = Kind("EOF")
	ENTER   = Kind("ENTER")

	IDENTIFIER = Kind("IDENTIFIER")
	NUMERAL    = Kind("NUMERAL")

	ASSIGN = Kind("=")

	// LET 예약어(keywords)
	LET = Kind("LET")
)

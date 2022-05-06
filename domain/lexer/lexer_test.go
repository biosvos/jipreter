package lexer

import (
	"github.com/stretchr/testify/assert"
	"jipreter/domain/lexer/token"
	"testing"
)

func TestLexer_Lexing(t *testing.T) {
	lexer := NewLexer("let variable = 3")

	tokens, err := lexer.Lexing()

	assert.NoError(t, err)
	for i, tt := range []struct {
		expectedToken *token.Token
	}{
		{
			expectedToken: token.NewToken(token.LET, "let"),
		}, {
			expectedToken: token.NewToken(token.IDENTIFIER, "variable"),
		}, {
			expectedToken: token.NewToken(token.ASSIGN, "="),
		}, {
			expectedToken: token.NewToken(token.NUMERAL, "3"),
		},
	} {
		assert.Equal(t, tt.expectedToken, tokens[i])
	}
}

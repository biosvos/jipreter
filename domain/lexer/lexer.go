package lexer

import (
	"jipreter/domain/lexer/token"
	"jipreter/domain/lexer/trie"
)

type Lexer struct {
	sourceCode string

	currentIndex int

	root *trie.Node
}

func NewLexer(sourceCode string) *Lexer {
	trieNode := trie.NewTrie("", 0)

	//trieNode.Register("=")
	return &Lexer{sourceCode: sourceCode}
}

//func (l *Lexer) registerTokenWork(s string) {
//
//}

func (l *Lexer) Lexing() ([]*token.Token, error) {
	var ret []*token.Token
	//
	//for l.more() {
	//	ch := l.next()
	//	log.Println(string(ch))
	//	var tok *token.Token
	//	switch ch {
	//	case ' ':
	//
	//	case '=':
	//		l.next()
	//		lit := l.drain()
	//		tok = token.NewToken(token.ASSIGN, lit)
	//	default:
	//
	//	}
	//}
	//for _, ch := range l.sourceCode {
	//
	//	switch ch {
	//	case '=':
	//		tok = token.NewToken(token.ASSIGN, string(ch))
	//	default:
	//		if isLetter(ch) {
	//
	//		}
	//	}
	//}

	return ret, nil
}

func isLetter(ch int32) bool {
	return false
}

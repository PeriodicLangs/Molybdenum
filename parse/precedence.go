package parse

import "github.com/westsi/molybdenum/lex"

const (
	_ = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

var precedences = map[lex.Token]int{
	lex.ASSIGN: EQUALS,
	lex.LT:     LESSGREATER,
	lex.GT:     LESSGREATER,
	lex.ADD:    SUM,
	lex.SUB:    SUM,
	lex.MUL:    PRODUCT,
	lex.DIV:    PRODUCT,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekTok.Tok]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curTok.Tok]; ok {
		return p
	}
	return LOWEST
}

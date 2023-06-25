package parse

import (
	"fmt"

	"github.com/westsi/molybdenum/lex"
)

type ParseReader struct {
	tokens []lex.LexedTok
	idx    int
	eof    bool
}

func NewParseReader(tokens []lex.LexedTok) *ParseReader {
	return &ParseReader{
		tokens: tokens,
		idx:    0,
		eof:    false,
	}
}

func (p *ParseReader) Read() lex.LexedTok {
	if p.eof {
		return lex.LexedTok{Pos: lex.Position{}, Tok: lex.EOF, Val: ""}
	}
	tok := p.tokens[p.idx]
	// fmt.Println(p.idx)
	p.idx++
	if p.idx >= len(p.tokens) {
		p.eof = true
	}
	return tok
}

func (p *ParseReader) PrintRem() {
	for _, tok := range p.tokens {
		fmt.Print(tok)
		fmt.Print(", ")
	}
}

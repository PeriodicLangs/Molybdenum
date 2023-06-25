package parse

import (
	"fmt"

	"github.com/westsi/molybdenum/ast"
	"github.com/westsi/molybdenum/lex"
)

type Parser struct {
	pr *ParseReader

	errors []string

	curTok  lex.LexedTok
	peekTok lex.LexedTok
}

func New(tokens []lex.LexedTok) *Parser {
	pr := NewParseReader(tokens)
	p := &Parser{pr: pr, errors: []string{}}
	p.nextTok()
	p.nextTok()
	return p
}

func (p *Parser) nextTok() {
	p.curTok = p.peekTok
	pt := p.pr.Read()
	p.peekTok = pt
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) Parse() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curTok.Tok != lex.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextTok()
	}

	return program
}

func (p *Parser) e(expected, actual lex.Token) {
	p.errors = append(p.errors, fmt.Sprintf("expected %s, got %s", expected, actual))
}

func (p *Parser) curTokenIs(t lex.Token) bool {
	return p.curTok.Tok == t
}
func (p *Parser) peekTokenIs(t lex.Token) bool {
	return p.peekTok.Tok == t
}
func (p *Parser) expectPeek(t lex.Token) bool {
	if p.peekTokenIs(t) {
		p.nextTok()
		return true
	} else {
		return false
	}
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curTok.Tok {
	case lex.RETURN:
		return p.parseReturn()
	case lex.VAR:
		return p.parseVar()
	default:
		return nil
	}
}

func (p *Parser) parseEDEF() *ast.EDEFStatement {
	stmt := &ast.EDEFStatement{Token: p.curTok}
	return stmt
}

func (p *Parser) parseVar() *ast.VarStatement {
	stmt := &ast.VarStatement{Token: p.curTok}

	if !p.expectPeek(lex.TYPEANNOT) {
		p.e(lex.TYPEANNOT, p.curTok.Tok)
	}
	stmt.Type = &ast.Type{Token: p.curTok, Value: p.curTok.Val}

	if !p.expectPeek(lex.IDENT) {
		p.e(lex.IDENT, p.curTok.Tok)
	}
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.String()}

	if !p.expectPeek(lex.ASSIGN) {
		p.e(lex.ASSIGN, p.curTok.Tok)
	}

	for !p.curTokenIs(lex.NEWLINE) {
		p.nextTok()
	}
	return stmt
}

func (p *Parser) parseReturn() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curTok}
	p.nextTok()

	for !p.curTokenIs(lex.NEWLINE) {
		p.nextTok()
	}

	return stmt
}

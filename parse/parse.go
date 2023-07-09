package parse

import (
	"fmt"
	"strconv"

	"github.com/westsi/molybdenum/ast"
	"github.com/westsi/molybdenum/lex"
)

type Parser struct {
	pr *ParseReader

	errors []string

	curTok  lex.LexedTok
	peekTok lex.LexedTok

	prefixParseFuncs map[lex.Token]prefixParseFunc
	infixParseFuncs  map[lex.Token]infixParseFunc
}

type (
	prefixParseFunc func() ast.Expression
	infixParseFunc  func(ast.Expression) ast.Expression
)

func (p *Parser) registerPrefix(tokenType lex.Token, fn prefixParseFunc) {
	p.prefixParseFuncs[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType lex.Token, fn infixParseFunc) {
	p.infixParseFuncs[tokenType] = fn
}
func (p *Parser) noPrefixParseFuncError(t lex.Token) {
	p.errors = append(p.errors, fmt.Sprintf("no prefix parse function for %s found", t))
}

func New(tokens []lex.LexedTok) *Parser {
	pr := NewParseReader(tokens)
	p := &Parser{pr: pr, errors: []string{}}
	p.nextTok()
	p.nextTok()

	p.prefixParseFuncs = make(map[lex.Token]prefixParseFunc)
	p.registerPrefix(lex.IDENT, p.parseIdentifier)
	p.registerPrefix(lex.INTLITERAL, p.parseIntegerLiteral)
	p.registerPrefix(lex.NOT, p.parsePrefixExpression)
	p.registerPrefix(lex.SUB, p.parsePrefixExpression)
	p.infixParseFuncs = make(map[lex.Token]infixParseFunc)
	p.registerInfix(lex.ADD, p.parseInfixExpression)
	p.registerInfix(lex.SUB, p.parseInfixExpression)
	p.registerInfix(lex.MUL, p.parseInfixExpression)
	p.registerInfix(lex.DIV, p.parseInfixExpression)
	p.registerInfix(lex.ASSIGN, p.parseInfixExpression)
	p.registerInfix(lex.LT, p.parseInfixExpression)
	p.registerInfix(lex.GT, p.parseInfixExpression)
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
		return p.parseReturnStatement()
	case lex.VAR:
		return p.parseVarStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curTok}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(lex.NEWLINE) {
		p.nextTok()
	}
	return stmt
}

func (p *Parser) parseExpression(prec int) ast.Expression {
	prefix := p.prefixParseFuncs[p.curTok.Tok]
	if prefix == nil {
		p.noPrefixParseFuncError(p.curTok.Tok)
		return nil
	}
	lExp := prefix()

	for !p.peekTokenIs(lex.NEWLINE) && prec < p.peekPrecedence() {
		infix := p.infixParseFuncs[p.peekTok.Tok]
		if infix == nil {
			return lExp
		}
		p.nextTok()
		lExp = infix(lExp)
	}

	return lExp
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curTok}
	val, err := strconv.ParseInt(p.curTok.Val, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer: error: %v", p.curTok.Val, err.Error()))
	}
	lit.Value = val
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	exp := &ast.PrefixExpression{
		Token:    p.curTok,
		Operator: p.curTok.Val,
	}
	p.nextTok()
	exp.Right = p.parseExpression(PREFIX)
	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	exp := &ast.InfixExpression{
		Token:    p.curTok,
		Operator: p.curTok.Val,
		Left:     left,
	}
	precedence := p.curPrecedence()
	p.nextTok()
	exp.Right = p.parseExpression(precedence)
	return exp
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
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
	p.nextTok()
	stmt.Value = p.parseExpressionStatement()
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curTok}
	p.nextTok()

	for !p.curTokenIs(lex.NEWLINE) {
		p.nextTok()
	}

	return stmt
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curTok, Value: p.curTok.String()}
}

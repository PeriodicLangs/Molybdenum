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
	p.registerPrefix(lex.TRUE, p.parseBoolean)
	p.registerPrefix(lex.FALSE, p.parseBoolean)
	p.registerPrefix(lex.IF, p.parseIfExpression)
	// p.registerPrefix(lex.FUNC, p.parseFunctionDefinition)
	// p.registerPrefix(lex.EFUNC, p.parseEntrypointFunctionDefinition)
	p.infixParseFuncs = make(map[lex.Token]infixParseFunc)
	p.registerInfix(lex.ADD, p.parseInfixExpression)
	p.registerInfix(lex.SUB, p.parseInfixExpression)
	p.registerInfix(lex.MUL, p.parseInfixExpression)
	p.registerInfix(lex.DIV, p.parseInfixExpression)
	p.registerInfix(lex.EQUALS, p.parseInfixExpression)
	p.registerInfix(lex.NOTEQUALS, p.parseInfixExpression)
	p.registerInfix(lex.LT, p.parseInfixExpression)
	p.registerInfix(lex.GT, p.parseInfixExpression)
	p.registerInfix(lex.LPAREN, p.parseCallExpression)
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
	// defer untrace(trace("parse"))
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
	// defer untrace(trace("parseStatement"))
	switch p.curTok.Tok {
	case lex.RETURN:
		return p.parseReturnStatement()
	case lex.VAR:
		stmt := p.parseVarStatement()
		return stmt
	case lex.NEWLINE:
		return nil
	case lex.FUNC:
		return p.parseFunctionDefinition()
	case lex.EFUNC:
		return p.parseEntrypointFunctionDefinition()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	// defer untrace(trace("parseExpressionStatement"))
	stmt := &ast.ExpressionStatement{Token: p.curTok}
	stmt.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(lex.NEWLINE) {
		p.nextTok()
	}
	return stmt
}

func (p *Parser) parseExpression(prec int) ast.Expression {
	// defer untrace(trace("parseExpression"))
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
	// defer untrace(trace("parseIntegerLiteral"))
	lit := &ast.IntegerLiteral{Token: p.curTok}
	val, err := strconv.ParseInt(p.curTok.Val, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Sprintf("could not parse %q as integer: error: %v", p.curTok.Val, err.Error()))
	}
	lit.Value = val
	return lit
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	// defer untrace(trace("parsePrefixExpression"))
	exp := &ast.PrefixExpression{
		Token:    p.curTok,
		Operator: p.curTok.Val,
	}
	p.nextTok()
	exp.Right = p.parseExpression(PREFIX)
	return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	// defer untrace(trace("parseInfixExpression"))
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

func (p *Parser) parseBoolean() ast.Expression {
	// defer untrace(trace("parseBoolean"))
	return &ast.Boolean{Token: p.curTok, Value: p.curTokenIs(lex.TRUE)}
}

func (p *Parser) parseIfExpression() ast.Expression {
	// defer untrace(trace("parseIfExpression"))
	exp := &ast.IfExpression{Token: p.curTok}

	if !p.expectPeek(lex.LPAREN) {
		return nil
	}
	p.nextTok()
	exp.Condition = p.parseExpression(LOWEST)
	if !p.expectPeek(lex.RPAREN) {
		return nil
	}

	if !p.expectPeek(lex.BLOCKSTART) {
		return nil
	}
	exp.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(lex.ELSE) {
		p.nextTok()

		if !p.expectPeek(lex.BLOCKSTART) {
			return nil
		}
		exp.Alternative = p.parseBlockStatement()
	}

	return exp
}
func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	// defer untrace(trace("parseBlockStatement"))
	block := &ast.BlockStatement{Token: p.curTok}
	block.Statements = []ast.Statement{}
	p.nextTok()
	if p.curTokenIs(lex.NEWLINE) {
		p.nextTok()
	}
	canCont := !p.curTokenIs(lex.BLOCKEND) && !p.curTokenIs(lex.EOF)
	for canCont {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		p.nextTok()
		if p.curTokenIs(lex.NEWLINE) { // blockend was consumed here because it was p.peekTokenIs which ignored the advance on line 243
			p.nextTok()
			p.nextTok()
		}
		canCont = !p.curTokenIs(lex.BLOCKEND) && !p.curTokenIs(lex.EOF)
	}
	return block
}

func (p *Parser) parseEntrypointFunctionDefinition() *ast.EntrypointFunctionDefinition {
	// defer untrace(trace("parseEntrypointFunctionDefinition"))
	efd := &ast.EntrypointFunctionDefinition{Token: p.curTok}
	p.nextTok()
	efd.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Val}
	if !p.expectPeek(lex.LPAREN) {
		return nil
	}
	if !p.expectPeek(lex.RPAREN) {
		return nil
	}
	// parameters do not get parsed because entrypoint functions should not have parameters
	if !p.expectPeek(lex.BLOCKSTART) {
		return nil
	}
	efd.Body = p.parseBlockStatement()
	return efd
}

func (p *Parser) parseFunctionDefinition() *ast.FunctionDefinition {
	// defer untrace(trace("parseFunctionDefinition"))
	fd := &ast.FunctionDefinition{Token: p.curTok}
	p.nextTok()
	fd.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Val}
	if !p.expectPeek(lex.LPAREN) {
		return nil
	}
	fd.Parameters = p.parseFunctionParameters()
	if !p.expectPeek(lex.BLOCKSTART) {
		return nil
	}
	fd.Body = p.parseBlockStatement()
	return fd
}

func (p *Parser) parseFunctionParameters() []*ast.Parameter {
	// defer untrace(trace("parseFunctionParameters"))
	parameters := []*ast.Parameter{}
	if p.peekTokenIs(lex.RPAREN) {
		p.nextTok()
		return parameters
	}
	p.nextTok()
	param := &ast.Parameter{}
	if !p.curTokenIs(lex.TYPEANNOT) {
		p.e(lex.TYPEANNOT, p.curTok.Tok)
	}
	param.Type = &ast.Type{Token: p.curTok, Value: p.curTok.Val}
	p.nextTok()
	if !p.curTokenIs(lex.IDENT) {
		p.e(lex.IDENT, p.curTok.Tok)
	}
	param.Token = p.curTok
	param.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Val}
	parameters = append(parameters, param)

	for p.peekTokenIs(lex.COMMA) {
		p.nextTok()
		p.nextTok()
		param := &ast.Parameter{}
		if !p.curTokenIs(lex.TYPEANNOT) {
			p.e(lex.TYPEANNOT, p.curTok.Tok)
		}
		param.Type = &ast.Type{Token: p.curTok, Value: p.curTok.Val}
		p.nextTok()
		if !p.curTokenIs(lex.IDENT) {
			p.e(lex.IDENT, p.curTok.Tok)
		}
		param.Token = p.curTok
		param.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Val}
		parameters = append(parameters, param)
	}
	if !p.expectPeek(lex.RPAREN) {
		return nil
	}
	return parameters
}

func (p *Parser) parseVarStatement() *ast.VarStatement {
	// defer untrace(trace("parseVarStatement"))
	stmt := &ast.VarStatement{Token: p.curTok}

	if !p.expectPeek(lex.TYPEANNOT) {
		p.e(lex.TYPEANNOT, p.curTok.Tok)
	}
	stmt.Type = &ast.Type{Token: p.curTok, Value: p.curTok.Val}

	if !p.expectPeek(lex.IDENT) {
		p.e(lex.IDENT, p.curTok.Tok)
	}
	stmt.Name = &ast.Identifier{Token: p.curTok, Value: p.curTok.Val}

	if !p.expectPeek(lex.ASSIGN) {
		p.e(lex.ASSIGN, p.curTok.Tok)
	}
	p.nextTok()
	stmt.Value = p.parseExpressionStatement()
	if p.peekTokenIs(lex.NEWLINE) {
		p.nextTok()
	}
	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	// defer untrace(trace("parseReturnStatement"))
	stmt := &ast.ReturnStatement{Token: p.curTok}
	p.nextTok()

	stmt.ReturnValue = p.parseExpression(LOWEST)
	if p.peekTokenIs(lex.NEWLINE) {
		p.nextTok()
	}

	return stmt
}

func (p *Parser) parseIdentifier() ast.Expression {
	// defer untrace(trace("parseIdentifier"))
	return &ast.Identifier{Token: p.curTok, Value: p.curTok.Val}
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	// defer untrace(trace("parseCallExpression"))
	exp := &ast.CallExpression{Token: p.curTok, Function: function}
	exp.Arguments = p.parseCallArguments()
	return exp
}

func (p *Parser) parseCallArguments() []ast.Expression {
	// defer untrace(trace("parseCallArguments"))
	args := []ast.Expression{}
	if p.peekTokenIs(lex.RPAREN) {
		p.nextTok()
		return args
	}
	p.nextTok()
	args = append(args, p.parseExpression(LOWEST))

	for p.peekTokenIs(lex.COMMA) {
		p.nextTok()
		p.nextTok()
		args = append(args, p.parseExpression(LOWEST))
	}
	if !p.expectPeek(lex.RPAREN) {
		return nil
	}
	return args
}

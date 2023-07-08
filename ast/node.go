package ast

import (
	"fmt"

	"github.com/westsi/molybdenum/lex"
)

type Node interface {
	Literal() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type EDEFStatement struct {
	Token lex.LexedTok
	Name  *Identifier
	Value *Block
}

func (es *EDEFStatement) statementNode() {}
func (es *EDEFStatement) Literal() string {
	return fmt.Sprintf("token: %s, name: %s, value: %s\n", es.Token.Tok.String(), es.Name.Literal(), es.Value.Literal())
}

type VarStatement struct {
	Token lex.LexedTok
	Name  *Identifier
	Value Expression
	Type  *Type
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) Literal() string {
	return fmt.Sprintf("token: %s, name: %s, value: %s, type: %s\n", vs.Token.Tok.String(), vs.Name.Literal(), vs.Value.Literal(), vs.Type.Literal())
	// return fmt.Sprintf("token: %s, name: %s, type: %s\n", vs.Token.Tok.String(), vs.Name.Literal(), vs.Type.Literal())
}

type Identifier struct {
	Token lex.LexedTok
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Literal() string {
	return fmt.Sprintf("{%s}", i.Value)
}

type Type struct {
	Token lex.LexedTok
	Value string
}

func (t *Type) statementNode() {}
func (t *Type) Literal() string {
	return fmt.Sprintf("{%s, %s}", t.Token.Tok.String(), t.Value)
}

type Block struct {
	Token      lex.LexedTok
	Statements []Statement
}

func (b *Block) statementNode() {}
func (b *Block) Literal() string {
	return fmt.Sprintf("token: %s, statements: %v\n", b.Token.String(), b.Statements)
}

type ReturnStatement struct {
	Token       lex.LexedTok
	ReturnValue Expression
}

func (ret *ReturnStatement) statementNode() {}
func (ret *ReturnStatement) Literal() string {
	return fmt.Sprintf("token: %s, value: %s\n", ret.Token.Tok.String(), ret.ReturnValue.Literal())
}

type ExpressionStatement struct {
	Token      lex.LexedTok
	Expression Expression
}

func (exp *ExpressionStatement) expressionNode() {}
func (exp *ExpressionStatement) statementNode()  {}
func (exp *ExpressionStatement) Literal() string {
	return fmt.Sprintf("token: %s, value: %s\n", exp.Token.Tok.String(), exp.Expression.Literal())
}

type IntegerLiteral struct {
	Token lex.LexedTok
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) Literal() string {
	return fmt.Sprintf("token: %s, value: %d\n", i.Token.Tok.String(), i.Value)
}

type PrefixExpression struct {
	Token    lex.LexedTok
	Operator string
	Right    Expression
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) Literal() string {
	return fmt.Sprintf("token: %s, operator: %s, right: %s\n", p.Token.Tok.String(), p.Operator, p.Right.Literal())
}

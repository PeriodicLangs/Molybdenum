package ast

import (
	"fmt"

	"github.com/westsi/molybdenum/lex"
)

type Node interface {
	Literal() string
	String() string
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
func (vs *VarStatement) String() string {
	return fmt.Sprintf("(%s, %s, = %s)", vs.Type.String(), vs.Name.String(), vs.Value.String())
}

type Identifier struct {
	Token lex.LexedTok
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) Literal() string {
	return fmt.Sprintf("{%s}", i.Value)
}
func (i *Identifier) String() string {
	return i.Value
}

type Type struct {
	Token lex.LexedTok
	Value string
}

func (t *Type) statementNode() {}
func (t *Type) Literal() string {
	return fmt.Sprintf("{%s, %s}", t.Token.Tok.String(), t.Value)
}
func (t *Type) String() string {
	return t.Value
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
func (ret *ReturnStatement) String() string {
	return fmt.Sprintf("(return %s)", ret.ReturnValue.String())
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
func (exp *ExpressionStatement) String() string {
	return exp.Expression.String()
}

type IntegerLiteral struct {
	Token lex.LexedTok
	Value int64
}

func (i *IntegerLiteral) expressionNode() {}
func (i *IntegerLiteral) Literal() string {
	return fmt.Sprintf("token: %s, value: %d\n", i.Token.Tok.String(), i.Value)
}
func (i *IntegerLiteral) String() string {
	return fmt.Sprintf("%d", i.Value)
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
func (p *PrefixExpression) String() string {
	return fmt.Sprintf("(%s %s)", p.Operator, p.Right.String())
}

type InfixExpression struct {
	Token    lex.LexedTok
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) Literal() string {
	return fmt.Sprintf("token: %s, left: %s, operator: %s, right: %s\n", i.Token.Tok.String(), i.Left.Literal(), i.Operator, i.Right.Literal())
}
func (i *InfixExpression) String() string {
	return fmt.Sprintf("(%s %s %s)", i.Left.String(), i.Operator, i.Right.String())
}

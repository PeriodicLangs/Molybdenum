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
	// Value Expression
	Type *Type
}

func (vs *VarStatement) statementNode() {}
func (vs *VarStatement) Literal() string {
	// return fmt.Sprintf("token: %s, name: %s, value: %s, type: %s\n", vs.Token.Tok.String(), vs.Name.Literal(), vs.Value.Literal(), vs.Type.Literal())
	return fmt.Sprintf("token: %s, name: %s, type: %s\n", vs.Token.Tok.String(), vs.Name.Literal(), vs.Type.Literal())
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

func (t *Type) expressionNode() {}
func (t *Type) Literal() string {
	return fmt.Sprintf("{%s, %s}", t.Token.Tok.String(), t.Value)
}

type Block struct {
	Token      lex.Token
	Statements []Statement
}

func (b *Block) statementNode() {}
func (b *Block) Literal() string {
	return fmt.Sprintf("token: %s, statements: %v\n", b.Token.String(), b.Statements)
}

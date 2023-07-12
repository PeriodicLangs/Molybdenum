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
	return i.Token.Val
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

type Boolean struct {
	Token lex.LexedTok
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) Literal() string {
	return fmt.Sprintf("token: %s, value: %t\n", b.Token.Tok.String(), b.Value)
}
func (b *Boolean) String() string {
	return fmt.Sprintf("%t", b.Value)
}

type IfExpression struct {
	Token       lex.LexedTok
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) expressionNode() {}
func (i *IfExpression) Literal() string {
	return fmt.Sprintf("token: %s, condition: %s, consequence: %s, alternative: %s\n", i.Token.Tok.String(), i.Condition.Literal(), i.Consequence.Literal(), i.Alternative.Literal())
}
func (i *IfExpression) String() string {
	if i.Alternative != nil {
		return fmt.Sprintf("(if %s %s else %s)", i.Condition.String(), i.Consequence.String(), i.Alternative.String())
	}
	return fmt.Sprintf("(if %s %s)", i.Condition.String(), i.Consequence.String())
}

type BlockStatement struct {
	Token      lex.LexedTok
	Statements []Statement
}

func (b *BlockStatement) statementNode() {}
func (b *BlockStatement) Literal() string {
	stmtString := ""
	for _, stmt := range b.Statements {
		stmtString += stmt.Literal()
		stmtString += ","
	}
	return fmt.Sprintf("token: %s, statements: %s\n", b.Token.Tok.String(), stmtString)
}
func (b *BlockStatement) String() string {
	s := ""
	for _, stmt := range b.Statements {
		s += "," + stmt.String()
	}
	return s
}

type FunctionDefinition struct {
	Token      lex.LexedTok
	Parameters []*Parameter
	Body       *BlockStatement
	Name       *Identifier
}

func (f *FunctionDefinition) expressionNode() {}
func (f *FunctionDefinition) Literal() string {
	return fmt.Sprintf("token: %s, parameters: %s, body: %s, name: %s\n", f.Token.Tok.String(), f.Parameters, f.Body.Literal(), f.Name.Literal())
}
func (f *FunctionDefinition) String() string {
	ps := ""
	for _, p := range f.Parameters {
		ps += p.String()
		ps += ", "
	}
	return fmt.Sprintf("(func %s (%s) {%s})", f.Name.String(), ps, f.Body.String())
}

type Parameter struct {
	Token lex.LexedTok
	Name  *Identifier
	Type  *Type
}

func (p *Parameter) expressionNode() {}
func (p *Parameter) Literal() string {
	return fmt.Sprintf("token: %s, name: %s, type: %s\n", p.Token.Tok.String(), p.Name.Literal(), p.Type.Literal())
}
func (p *Parameter) String() string {
	return fmt.Sprintf("%s %s", p.Name.String(), p.Type.String())
}

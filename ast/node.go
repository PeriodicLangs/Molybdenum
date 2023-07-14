package ast

import (
	"fmt"
	"strings"

	"github.com/westsi/molybdenum/lex"
)

type Node interface {
	Literal() string
	String() string
}

type Statement interface {
	Node
	statementNode()
	NType() string
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
func (vs *VarStatement) NType() string  { return "VarStatement" }
func (vs *VarStatement) Literal() string {
	return fmt.Sprintf("token: %s, name: %s, value: %s, type: %s\n", vs.Token.Tok.String(), vs.Name.Literal(), vs.Value.Literal(), vs.Type.Literal())
	// return fmt.Sprintf("token: %s, name: %s, type: %s\n", vs.Token.Tok.String(), vs.Name.Literal(), vs.Type.Literal())
}
func (vs *VarStatement) String() string {
	return fmt.Sprintf("(%s %s = %s)", vs.Type.String(), vs.Name.String(), vs.Value.String())
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

func (t *Type) expressionNode() {}
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
func (ret *ReturnStatement) NType() string  { return "ReturnStatement" }
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
func (exp *ExpressionStatement) NType() string   { return "ExpressionStatement" }
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
func (b *BlockStatement) NType() string  { return "BlockStatement" }
func (b *BlockStatement) Literal() string {
	stmtString := []string{}
	for _, stmt := range b.Statements {
		stmtString = append(stmtString, stmt.Literal())
	}
	return fmt.Sprintf("token: %s, statements: %s\n", b.Token.Tok.String(), strings.Join(stmtString, ", "))
}
func (b *BlockStatement) String() string {
	s := []string{}
	for _, stmt := range b.Statements {
		s = append(s, stmt.String())
	}
	return strings.Join(s, ", ")
}

type FunctionDefinition struct {
	Token      lex.LexedTok
	Parameters []*Parameter
	Body       *BlockStatement
	Name       *Identifier
}

func (f *FunctionDefinition) expressionNode() {}
func (f *FunctionDefinition) statementNode()  {}
func (f *FunctionDefinition) NType() string   { return "FunctionDefinition" }
func (f *FunctionDefinition) Literal() string {
	return fmt.Sprintf("token: %s, parameters: %s, body: %s, name: %s\n", f.Token.Tok.String(), f.Parameters, f.Body.Literal(), f.Name.Literal())
}
func (f *FunctionDefinition) String() string {
	ps := []string{}
	for _, p := range f.Parameters {
		ps = append(ps, p.String())
	}
	return fmt.Sprintf("(func %s (%s) {%s})", f.Name.String(), strings.Join(ps, ", "), f.Body.String())
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

type CallExpression struct {
	Token     lex.LexedTok
	Function  Expression
	Arguments []Expression
}

func (c *CallExpression) expressionNode() {}
func (c *CallExpression) Literal() string {
	return fmt.Sprintf("token: %s, function: %s, arguments: %s\n", c.Token.Tok.String(), c.Function.Literal(), c.Arguments)
}
func (c *CallExpression) String() string {
	args := []string{}
	for _, arg := range c.Arguments {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("(%s(%s))", c.Function.String(), strings.Join(args, ", "))
}

type EntrypointFunctionDefinition struct {
	Token lex.LexedTok
	// Entrypoint functions have no parameters
	// In addition, they should never be called by the user - calls for them are handled by the compiler
	Body *BlockStatement
	Name *Identifier
}

func (e *EntrypointFunctionDefinition) expressionNode() {}
func (e *EntrypointFunctionDefinition) statementNode()  {}
func (e *EntrypointFunctionDefinition) NType() string   { return "EntrypointFunctionDefinition" }
func (e *EntrypointFunctionDefinition) Literal() string {
	return fmt.Sprintf("token: %s, body: %s, name: %s\n", e.Token.Tok.String(), e.Body.Literal(), e.Name.Literal())
}
func (e *EntrypointFunctionDefinition) String() string {
	return fmt.Sprintf("(entrypoint %s {%s})", e.Name.String(), e.Body.String())
}

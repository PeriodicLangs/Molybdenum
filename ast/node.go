package ast

import "fmt"

type Node interface {
	String() string
	getValue() string
	getType() string
	getChildren() []Node
}

type IdentNode struct {
	Type     string
	Value    string // in IdentNode, holds name of variable
	Children []Node
}

func (i IdentNode) getValue() string    { return i.Value }
func (i IdentNode) getType() string     { return i.Type }
func (i IdentNode) getChildren() []Node { return i.Children }

func (i IdentNode) String() string {
	s := fmt.Sprintf("%s, %s\n", i.Type, i.Value)
	for _, child := range i.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

type KeywordNode struct {
	Type     string
	Value    string // in KeywordNode, holds keyword
	Children []Node
}

func (k KeywordNode) getValue() string    { return k.Value }
func (k KeywordNode) getType() string     { return k.Type }
func (k KeywordNode) getChildren() []Node { return k.Children }

func (k KeywordNode) String() string {
	s := fmt.Sprintf("%s, %s\n", k.Type, k.Value)
	for _, child := range k.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

// type TypeAnnotationNode struct {
// 	Type     string
// 	Children []Node
// }

// func (t *TypeAnnotationNode) getValue() string    { return t.Type }
// func (t *TypeAnnotationNode) getType() string     { return t.Type }
// func (t *TypeAnnotationNode) getChildren() []Node { return t.Children }
// func (t *TypeAnnotationNode) String() string {
// 	s := fmt.Sprintf("%s\n", t.Type)
// 	for _, child := range t.Children {
// 		s = s + "  - " + child.String() + "\n"
// 	}
// 	return s
// }

// type OperatorNode struct {
// 	Type     string
// 	Children []Node
// 	Value    string // in OperatorNode, holds operator - one of =+-*/%
// }

// func (o *OperatorNode) getValue() string    { return o.Value }
// func (o *OperatorNode) getType() string     { return o.Type }
// func (o *OperatorNode) getChildren() []Node { return o.Children }
// func (o *OperatorNode) String() string {
// 	s := fmt.Sprintf("%s\n", o.Type)
// 	for _, child := range o.Children {
// 		s = s + "  - " + child.String() + "\n"
// 	}
// 	return s
// }

type BlockNode struct {
	Type     string
	Children []Node
}

func (b BlockNode) getValue() string    { return b.Type }
func (b BlockNode) getType() string     { return b.Type }
func (b BlockNode) getChildren() []Node { return b.Children }

func (b BlockNode) String() string {
	s := fmt.Sprintf("%s\n", b.Type)
	for _, child := range b.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

type StringLiteralNode struct {
	Type     string
	Value    string // holds string literal
	Children []Node
}

func (s StringLiteralNode) getValue() string    { return s.Value }
func (s StringLiteralNode) getType() string     { return s.Type }
func (s StringLiteralNode) getChildren() []Node { return s.Children }

func (sn StringLiteralNode) String() string {
	s := fmt.Sprintf("%s\n", sn.Type)
	for _, child := range sn.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

// type IntLiteralNode struct {
// 	Type     string
// 	Value    string // holds int literal
// 	Children []Node
// }

// func (i *IntLiteralNode) getValue() string    { return i.Value }
// func (i *IntLiteralNode) getType() string     { return i.Type }
// func (i *IntLiteralNode) getChildren() []Node { return i.Children }
// func (i *IntLiteralNode) String() string {
// 	s := fmt.Sprintf("%s\n", i.Type)
// 	for _, child := range i.Children {
// 		s = s + "  - " + child.String() + "\n"
// 	}
// 	return s
// }

type EntryPointDeclNode struct {
	Type     string // hold name of entry point - one of validEntryPointNames
	Children []Node
	// arguments are not needed because by language design, an entry point has no arguments
}

func (e EntryPointDeclNode) getValue() string    { return e.Type }
func (e EntryPointDeclNode) getType() string     { return e.Type }
func (e EntryPointDeclNode) getChildren() []Node { return e.Children }

func (e EntryPointDeclNode) String() string {
	s := fmt.Sprintf("%s\n", e.Type)
	for _, child := range e.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

type FuncCallNode struct {
	Value    string
	Children []Node
}

func (f FuncCallNode) getValue() string    { return f.Value }
func (f FuncCallNode) getType() string     { return "FuncCallNode" }
func (f FuncCallNode) getChildren() []Node { return f.Children }

func (f FuncCallNode) String() string {
	s := "FuncCallNode"
	for _, child := range f.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

type ArgsNode struct {
	Children []Node
}

func (a ArgsNode) getValue() string    { return "ArgsNode" }
func (a ArgsNode) getType() string     { return "" }
func (a ArgsNode) getChildren() []Node { return a.Children }

func (a ArgsNode) String() string {
	s := "ArgsNode"
	for _, child := range a.Children {
		s = s + "  - " + child.String() + "\n"
	}
	return s
}

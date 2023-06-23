package main

type AST struct {
	Children []Node
}

func (a *AST) String() string {
	s := ""
	for _, child := range a.Children {
		s = s + child.String() + "\n"
	}
	return s
}

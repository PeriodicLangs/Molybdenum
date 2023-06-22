package main

import (
	"fmt"
)

func Parse(tokens []*LexedTok) (*AST, error) {
	// look at how https://github.com/shafinsiddique/page/blob/master/page/rdParser.go is structured
	var ast AST
	for _, lt := range tokens {
		node, e := parseNode(lt)
		if e != nil {
			return nil, e
		}
		ast.Children = append(ast.Children, node)
	}
	return &ast, nil
}

func pop(a []*LexedTok) *LexedTok {
	return a[len(a)-1]
}

func peek(a []*LexedTok, i int) *LexedTok {
	return a[0]
}

func parseNode(lt *LexedTok) (Node, error) {
	switch lt.tok {
	case KEYWORD:
		return parseKeyword()
	case IDENT:
		return parseIdent()
	case TYPEANNOT:
		return parseTypeAnnotation()
	case BLOCKSTART:
		return parseBlock()
	}
	return nil, fmt.Errorf("unexpected token: %v", lt.tok)
}

func parseKeyword() (*KeywordNode, error) {
	return nil, nil
}

func parseIdent() (*IdentNode, error) {
	return nil, nil
}

func parseTypeAnnotation() (*TypeAnnotationNode, error) {
	return nil, nil
}

func parseBlock() (*BlockNode, error) {
	return nil, nil
}

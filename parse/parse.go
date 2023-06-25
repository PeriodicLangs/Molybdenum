package parse

import (
	"fmt"

	"github.com/westsi/molybdenum/ast"
	"github.com/westsi/molybdenum/lex"
)

var pr *ParseReader

func Parse(tokens []*lex.LexedTok) (*ast.AST, error) {
	var ast ast.AST

	return &ast, nil
}

func expect(expected, actual lex.Token) error {
	if expected == actual {
		return nil
	} else {
		return fmt.Errorf("expected %v, got %v", expected, actual)
	}
}

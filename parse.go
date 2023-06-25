package main

import (
	"fmt"
)

var pr *ParseReader

func Parse(tokens []*LexedTok) (*AST, error) {
	var ast AST

	return &ast, nil
}

func expect(expected, actual Token) error {
	if expected == actual {
		return nil
	} else {
		return fmt.Errorf("expected %v, got %v", expected, actual)
	}
}

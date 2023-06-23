package main

import (
	"fmt"
	"io"
)

var pr *ParseReader

func Parse(tokens []*LexedTok) (*AST, error) {
	pr = NewParseReader(tokens)
	var ast AST
	var err error = nil
	for err != io.EOF {
		fmt.Println("looping")
		readtok, err := pr.Read()
		if err == io.EOF {
			return &ast, nil
		} else if err != nil {
			return nil, err
		}
		node, parseErr := parseNode(readtok)
		if parseErr != nil {
			return nil, parseErr
		}
		if node != nil {
			ast.Children = append(ast.Children, node)
		}
	}
	return &ast, nil
}

func expect(expected *LexedTok, actual *LexedTok) error {
	if expected.Tok == actual.Tok {
		return nil
	} else {
		return fmt.Errorf("expected %v, got %v", expected.Tok, actual.Tok)
	}
}

func parseNode(lt *LexedTok) (Node, error) {
	switch lt.Tok {
	case KEYWORD:
		return parseKeyword(lt)
	case IDENT:
		return parseIdent(lt)
	case NEWLINE:
		return nil, nil
	case BLOCKSTART:
		return parseBlock(lt)
	case EOF:
		return nil, nil
	}
	return nil, fmt.Errorf("unexpected token: %v", lt.Tok)
}

func parseKeyword(tok *LexedTok) (Node, error) {
	kw := tok.Val
	if kw == "edef" {
		// parse next IDENT
		t, err := pr.Read()
		if err != nil {
			return nil, err
		}
		ident, err := parseIdent(t)
		if err != nil {
			return nil, err
		}
		pr.Read()
		pr.Read()
		edn := EntryPointDeclNode{
			Type:     "ENTRYPOINTDECLNODE",
			Children: []Node{ident},
		}
		return edn, nil
	}
	return nil, fmt.Errorf("illegal keyword %s", kw)
}

func parseIdent(tok *LexedTok) (Node, error) {
	// if ident followed by LPAREN then:
	// if ident preceded by KEYWORD (edef, def) is FuncDecl
	// if not is FuncCall
	return IdentNode{
		Type:  "IDENTNODE",
		Value: tok.Val,
	}, nil
}

func parseBlock(tok *LexedTok) (Node, error) {
	return nil, nil
}

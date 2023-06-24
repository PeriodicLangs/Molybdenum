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

func expect(expected, actual Token) error {
	if expected == actual {
		return nil
	} else {
		return fmt.Errorf("expected %v, got %v", expected, actual)
	}
}

func parseNode(lt *LexedTok) (Node, error) {
	switch lt.Tok {
	case KEYWORD:
		return parseKeyword(lt)
	case IDENT:
		return parseIdent(nil, lt)
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
		ident, err := parseIdent(tok, t)
		if err != nil {
			return nil, err
		}
		edn := EntryPointDeclNode{
			Type:     "ENTRYPOINTDECLNODE",
			Children: []Node{ident},
		}
		return edn, nil
	}
	return nil, fmt.Errorf("illegal keyword %s", kw)
}

func parseIdent(prec, ident *LexedTok) (Node, error) {
	// if ident followed by LPAREN then:
	// if ident preceded by KEYWORD (edef, def) is FuncDecl
	// if not is FuncCall
	t, err := pr.Read()
	if err != nil {
		return nil, err
	}
	if t.Tok == LPAREN {
		if prec == nil {
			pr.Read()
			return FuncCallNode{
				Value:    ident.Val,
				Children: []Node{},
			}, nil
		} else if prec.Tok == KEYWORD {
			if prec.Val == "edef" {
				_t, err := pr.Read() // this is to consume RPAREN of edef'ed function
				if err != nil {
					return nil, err
				}
				err = expect(RPAREN, _t.Tok)
				if err != nil {
					return nil, err
				}

				_t, err = pr.Read() // this is to consume BLOCKSTART of edef'ed function
				if err != nil {
					return nil, err
				}
				err = expect(BLOCKSTART, _t.Tok)
				if err != nil {
					return nil, err
				} else {
					b, err := parseBlock(_t)
					if err != nil {
						return nil, err
					}
					return IdentNode{
						Type:     "IDENTNODE",
						Value:    ident.Val,
						Children: []Node{b},
					}, nil
				}
			}
		} else {
			return nil, fmt.Errorf("parseIdent - unexpected token: %v", t.Tok)
		}
	}
	return IdentNode{
		Type:  "IDENTNODE",
		Value: ident.Val,
	}, nil
}

func parseBlock(tok *LexedTok) (Node, error) {
	t, err := pr.Read()
	if err != nil {
		return nil, err
	}
	var children []Node
	for t.Tok != BLOCKEND {
		t, err = pr.Read()
		if err != nil {
			return nil, err
		}
		if t.Tok == BLOCKEND {
			break
		}
		c, err := parseNode(t)
		if err != nil {
			return nil, err
		}
		children = append(children, c)
	}
	_, err = pr.Read()
	if err != nil {
		return nil, err
	}
	pr.PrintRem()
	return BlockNode{
		Type:     "BLOCKNODE",
		Children: children,
	}, nil
}

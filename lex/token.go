package lex

import "fmt"

type Token int

const (
	EOF Token = iota
	ILLEGAL
	IDENT
	// language keywords
	EFUNC
	FUNC
	METH
	VAR
	IF
	ELSE
	FOR
	WHILE
	RETURN
	BREAK
	CONTINUE
	AS
	// end of language keywords
	TYPEANNOT
	IMPORT
	ASSIGN
	ADD
	MUL
	SUB
	DIV
	MOD
	LPAREN
	RPAREN
	LSQRBRAC
	RSQRBRAC
	BLOCKSTART
	BLOCKEND
	INTLITERAL
	STRINGLITERAL
	DOT
	NEWLINE
	AND
	NOT
	GT
	LT
	TRUE
	FALSE
	NOTEQUALS
	EQUALS
	COMMA
)

var tokens = []string{
	EOF:           "EOF",
	ILLEGAL:       "ILLEGAL",
	IDENT:         "IDENT",
	EFUNC:         "EFUNC",
	FUNC:          "FUNC",
	METH:          "METH",
	VAR:           "VAR",
	IF:            "IF",
	ELSE:          "ELSE",
	FOR:           "FOR",
	WHILE:         "WHILE",
	RETURN:        "RETURN",
	BREAK:         "BREAK",
	CONTINUE:      "CONTINUE",
	AS:            "AS",
	TYPEANNOT:     "TYPEANNOT",
	IMPORT:        "IMPORT", // right now import just exists, has no functionality yet
	ASSIGN:        "ASSIGN",
	ADD:           "ADD",
	MUL:           "MUL",
	SUB:           "SUB",
	DIV:           "DIV",
	MOD:           "MOD",
	LPAREN:        "LPAREN",
	RPAREN:        "RPAREN",
	LSQRBRAC:      "LSQRBRAC",
	RSQRBRAC:      "RSQRBRAC",
	BLOCKSTART:    "BLOCKSTART",
	BLOCKEND:      "BLOCKEND",
	STRINGLITERAL: "STRINGLITERAL",
	INTLITERAL:    "INTLITERAL",
	DOT:           "DOT",
	NEWLINE:       "NEWLINE",
	AND:           "AND",
	NOT:           "NOT",
	GT:            "GT",
	LT:            "LT",
	TRUE:          "TRUE",
	FALSE:         "FALSE",
	NOTEQUALS:     "NOTEQUALS",
	EQUALS:        "EQUALS",
	COMMA:         "COMMA",
}

var keywords = []string{
	"efunc",
	"func",
	"meth",
	"var",
	"if",
	"else",
	"for",
	"while",
	"return",
	"break",
	"continue",
	"as",
}

var kwmap = map[string]Token{
	"efunc":    EFUNC,
	"func":     FUNC,
	"meth":     METH,
	"var":      VAR,
	"if":       IF,
	"else":     ELSE,
	"for":      FOR,
	"while":    WHILE,
	"return":   RETURN,
	"break":    BREAK,
	"continue": CONTINUE,
	"as":       AS,
}

var types = []string{
	"string",
	"int",
	"float",
	"double",
	"bool",
}

func (t Token) String() string {
	return tokens[t]
}

type LexedTok struct {
	Pos Position
	Tok Token
	Val string
}

func NewLexedTok(pos Position, tok Token, val string) LexedTok {
	return LexedTok{
		Pos: pos,
		Tok: tok,
		Val: val,
	}
}

func (lt *LexedTok) String() string {
	return fmt.Sprint(lt.Pos) + " " + lt.Tok.String() + " " + lt.Val
}

var datatypes = map[Token]string{
	INTLITERAL:    "int",
	STRINGLITERAL: "string",
	TRUE:          "bool",
	FALSE:         "bool",
}

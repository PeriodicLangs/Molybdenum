package main

type Token int

const (
	EOF Token = iota
	ILLEGAL
	IDENT
	KEYWORD
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
)

var tokens = []string{
	EOF:           "EOF",
	ILLEGAL:       "ILLEGAL",
	IDENT:         "IDENT",
	KEYWORD:       "KEYWORD",
	TYPEANNOT:     "TYPEANNOT", // TODO
	IMPORT:        "IMPORT",    // TODO
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
}

var keywords = []string{
	"edef",
	"def",
	"mdef",
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

func (t Token) String() string {
	return tokens[t]
}

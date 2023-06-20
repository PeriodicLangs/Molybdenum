package main

type Token int

const (
	EOF Token = iota
	ILLEGAL
	IDENT
	INT
	SEMI

	ADD
	SUB
	MUL
	DIV

	ASSIGN
)

var tokens = []string{
	EOF:     "EOF",
	ILLEGAL: "ILLEGAL",
	IDENT:   "IDENT",
	INT:     "INT",
	SEMI:    ";",
	ADD:     "+",
	SUB:     "-",
	MUL:     "*",
	DIV:     "/",
	ASSIGN:  "=",
}

func (t Token) String() string {
	return tokens[t]
}

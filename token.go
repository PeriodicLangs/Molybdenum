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
	TYPEANNOT:     "TYPEANNOT",
	IMPORT:        "IMPORT", // right now import just exists, has no functionality yet
	ASSIGN:        "=",
	ADD:           "+",
	MUL:           "*",
	SUB:           "-",
	DIV:           "/",
	MOD:           "%",
	LPAREN:        "(",
	RPAREN:        ")",
	LSQRBRAC:      "[",
	RSQRBRAC:      "]",
	BLOCKSTART:    "{",
	BLOCKEND:      "}",
	STRINGLITERAL: "STRINGLITERAL",
	INTLITERAL:    "INTLITERAL",
	DOT:           ".",
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
	pos Position
	tok Token
	val string
}

func NewLexedTok(pos Position, tok Token, val string) *LexedTok {
	return &LexedTok{
		pos: pos,
		tok: tok,
		val: val,
	}
}

var validEntryPointNames = []string{
	"main",
	// "init", ADD LATER!!!
	// "initOnce", ADD LATER!!!
}

package main

import (
	"fmt"
	"os"
)

func main() {
	// reader := strings.NewReader("edef main() {\n" +
	// 	"Print(\"Hello World\")\n" +
	// 	"}")

	reader, err := os.Open("basic.mn")
	if err != nil {
		panic(err)
	}
	lexer := NewLexer(reader)
	var tokens []*LexedTok
	pos, tok, val := lexer.Lex()
	tokens = append(tokens, NewLexedTok(pos, tok, val))
	for tok != EOF {
		// fmt.Println(pos, tok, val)
		pos, tok, val = lexer.Lex()
		tokens = append(tokens, NewLexedTok(pos, tok, val))
	}
	// fmt.Println(pos, tok, val)
	for _, lt := range tokens {
		fmt.Println(lt.pos, lt.tok, lt.val)
	}
	// tokens := []*LexedTok{NewLexedTok(Position{1, 1}, IDENT, "main"), NewLexedTok(Position{2, 1}, BLOCKSTART, "{"), NewLexedTok(Position{3, 1}, TYPEANNOT, "string"), NewLexedTok(Position{3, 6}, IDENT, "s"), NewLexedTok(Position{3, 8}, ASSIGN, "="), NewLexedTok(Position{3, 10}, STRINGLITERAL, "Hello World!"), NewLexedTok(Position{4, 1}, BLOCKEND, "}"), NewLexedTok(Position{5, 1}, EOF, "")}
	// _ = Parse(tokens)
}

package main

import (
	"fmt"
	"os"

	"github.com/westsi/molybdenum/lex"
	"github.com/westsi/molybdenum/parse"
)

func main() {
	// reader := strings.NewReader("edef main() {\n" +
	// 	"Print(\"Hello World\")\n" +
	// 	"}")

	reader, err := os.Open("molybdenum/basic.mn")
	if err != nil {
		panic(err)
	}
	lexer := lex.NewLexer(reader)
	var tokens []*lex.LexedTok
	pos, tok, val := lexer.Lex()
	tokens = append(tokens, lex.NewLexedTok(pos, tok, val))
	for tok != lex.EOF {
		// fmt.Println(pos, tok, val)
		pos, tok, val = lexer.Lex()
		tokens = append(tokens, lex.NewLexedTok(pos, tok, val))
	}
	// fmt.Println(pos, tok, val)
	// for _, lt := range tokens {
	// 	fmt.Println(lt.Pos, lt.Tok, lt.Val)
	// }
	// tokens := []*LexedTok{NewLexedTok(Position{1, 1}, IDENT, "main"), NewLexedTok(Position{2, 1}, BLOCKSTART, "{"), NewLexedTok(Position{3, 1}, TYPEANNOT, "string"), NewLexedTok(Position{3, 6}, IDENT, "s"), NewLexedTok(Position{3, 8}, ASSIGN, "="), NewLexedTok(Position{3, 10}, STRINGLITERAL, "Hello World!"), NewLexedTok(Position{4, 1}, BLOCKEND, "}"), NewLexedTok(Position{5, 1}, EOF, "")}
	ast, err := parse.Parse(tokens)
	if err != nil {
		panic("ERROR: " + err.Error())
	}
	fmt.Println(ast)
}

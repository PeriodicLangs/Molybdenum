package main

import (
	"fmt"
	"os"
)

func main() {
	// reader := strings.NewReader("edef main() {\n" +
	// 	"Print(\"Hello World\")\n" +
	// 	"}")

	reader, err := os.Open("src/main.mn")
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
}

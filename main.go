package main

import (
	"fmt"
	"strings"
)

func main() {
	reader := strings.NewReader("edef main() {\n" +
		"Print(\"Hello World\")\n" +
		"}")
	lexer := NewLexer(reader)
	pos, tok, val := lexer.Lex()
	for tok != EOF {
		fmt.Println(pos, tok, val)
		pos, tok, val = lexer.Lex()
	}
	fmt.Println(pos, tok, val)
}

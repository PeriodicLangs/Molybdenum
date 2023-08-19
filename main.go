package main

import (
	"fmt"
	"os"

	"github.com/westsi/molybdenum/lex"
	"github.com/westsi/molybdenum/parse"
)

func main() {
	// reader := strings.NewReader("efunc main() {\n" +
	// 	"Print(\"Hello World\")\n" +
	// 	"}")

	// reader, err := os.Open("molybdenum/verify_system_test.mn")
	reader, err := os.Open("molybdenum/basic.mn")
	if err != nil {
		panic(err)
	}
	lexer := lex.NewLexer(reader)
	var tokens []lex.LexedTok
	pos, tok, val := lexer.Lex()
	tokens = append(tokens, lex.NewLexedTok(pos, tok, val))
	for tok != lex.EOF {
		pos, tok, val = lexer.Lex()
		tokens = append(tokens, lex.NewLexedTok(pos, tok, val))
	}

	for _, tok := range tokens {
		fmt.Println(tok)
	}

	p := parse.New(tokens)
	ast := p.Parse()
	fmt.Printf("Errors: %s\n", p.Errors())
	fmt.Println(ast.String())
}

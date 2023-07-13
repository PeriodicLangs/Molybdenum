package main

import (
	"fmt"
	"os"

	"github.com/westsi/molybdenum/parse"

	"github.com/westsi/molybdenum/codegen"
	"github.com/westsi/molybdenum/lex"
)

func main() {
	// reader := strings.NewReader("efunc main() {\n" +
	// 	"Print(\"Hello World\")\n" +
	// 	"}")

	reader, err := os.Open("molybdenum/functions.mn")
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

	codegen.Gen(ast)
}

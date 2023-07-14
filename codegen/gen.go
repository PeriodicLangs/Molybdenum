package codegen

import (
	"os"
	"os/exec"

	"github.com/llir/llvm/ir"
	"github.com/westsi/molybdenum/ast"
)

var m *ir.Module

func Gen(program *ast.Program) {
	m = ir.NewModule()
	for _, stmt := range program.Statements {
		// stmt.Gen()
		stmt.Literal()
	}
	compile()
}

func compile() {
	println(m.String())
	f, _ := os.Create("ir.ll")
	defer f.Close()
	m.WriteTo(f)

	cmd := exec.Command("clang", "ir.ll")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

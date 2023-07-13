package codegen

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/llir/llvm/ir"
	"github.com/westsi/molybdenum/ast"
)

var m *ir.Module

func Gen(program *ast.Program) {
	m = ir.NewModule()
	for _, stmt := range program.Statements {
		generateStatement(stmt)
	}
	compile()
}

func generateStatement(stmt ast.Statement) {
	// might not work
	switch stmt := stmt.(type) {
	case *ast.ExpressionStatement:
		generateExpressionStatement(stmt)
	case *ast.BlockStatement:
		generateBlockStatement(stmt)
	case *ast.ReturnStatement:
		generateReturnStatement(stmt)
	}
}

func generateExpressionStatement(stmt *ast.ExpressionStatement) {
	fmt.Println("Generating expression statement")
}

func generateBlockStatement(stmt *ast.BlockStatement) {
	fmt.Println("Generating block statement")
}

func generateReturnStatement(stmt *ast.ReturnStatement) {
	fmt.Println("Generating return statement")
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

package verify

import (
	"fmt"

	"github.com/westsi/molybdenum/ast"
)

// Verifies that all referenced variables/functions are defined

// Does not modify AST

/*
Module checks
	- variables are defined whenever used !
	- functions are defined - this can be after use !
	- variable values are always of the variable type when assigned !
	- the correct number and type of parameters are passed to functions !
	- the correct (number) and type of return value(s) are returned FROM ALL BRANCHES !
	- functions are not overloaded or redefined - TODO - add to verify_system_test.mn
	- variables are not redefined !
*/

var errors []error
var symtab SymTab
var fcallchecks []string // holds names of called functions that have not yet been defined

func Verify(prog ast.Program) []error {
	errors = []error{}
	symtab = SymTab{make(map[string]SymTabEntry)}
	fcallchecks = []string{}

	for _, stmt := range prog.Statements {
		fmt.Println("looping")
		switch st := stmt.(type) {
		case *ast.VarStatement:
			errors = append(errors, VerifyVarStatement(st)...)
		case *ast.FunctionDefinition:
			errors = append(errors, VerifyFunctionDefinition(st)...)
		case *ast.BlockStatement:
			errors = append(errors, VerifyBlockStatement(st)...)
		case *ast.EntrypointFunctionDefinition:
			errors = append(errors, VerifyEntrypointFunctionDefinition(st)...)
		case *ast.ExpressionStatement:
			errors = append(errors, VerifyExpressionStatement(st)...)
		case *ast.ReturnStatement:
			errors = append(errors, VerifyReturnStatement(st)...)
		default:
			errors = append(errors, fmt.Errorf("unknown statement type: %T", stmt))
		}
	}

	return errors
}

func VerifyVarStatement(vs *ast.VarStatement) []error {
	e := []error{}
	name := vs.Name.Value
	entry := symtab.Entries[name]
	fmt.Println(entry)
	fmt.Println(symtab)
	// check that this is the first definition of this identifier
	if entry != nil {
		e = append(e, fmt.Errorf("already defined %s", name))
	} else {
		symtab.Entries[name] = SymTabVariable{Type: vs.Type.Value}
	}
	// check that the value assigned to it is of the correct type
	if getLiteralType(vs.Value.(*ast.ExpressionStatement).Expression) != vs.Type.Value {
		// this cast to *ast.ExpressionStatement will never fail because the current parsing setup means that vs.Value is always a *ast.ExpressionStatement
		e = append(e, fmt.Errorf("type mismatch: expected %s, got %s", vs.Type.Value, getLiteralType(vs.Value.(*ast.ExpressionStatement).Expression)))
	}

	return e
}

func VerifyFunctionDefinition(f *ast.FunctionDefinition) []error { e := []error{}; return e }
func VerifyBlockStatement(f *ast.BlockStatement) []error         { e := []error{}; return e }
func VerifyEntrypointFunctionDefinition(f *ast.EntrypointFunctionDefinition) []error {
	e := []error{}
	return e
}
func VerifyExpressionStatement(f *ast.ExpressionStatement) []error { e := []error{}; return e }
func VerifyReturnStatement(f *ast.ReturnStatement) []error         { e := []error{}; return e }

func getLiteralType(lit ast.Expression) string {
	switch lit.(type) {
	case *ast.IntegerLiteral:
		return "int"
	// case *ast.StringLiteral:
	// 	return "string"
	case *ast.Boolean:
		return "bool"
	default:
		return "unknown"
	}
}

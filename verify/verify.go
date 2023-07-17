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
		verifyStatement(stmt)
	}

	return errors
}

func verifyStatement(stmt ast.Statement) {
	switch st := stmt.(type) {
	case *ast.VarStatement:
		errors = append(errors, verifyVarStatement(st)...)
	case *ast.FunctionDefinition:
		errors = append(errors, verifyFunctionDefinition(st)...)
	case *ast.BlockStatement:
		errors = append(errors, verifyBlockStatement(st)...)
	case *ast.EntrypointFunctionDefinition:
		errors = append(errors, verifyEntrypointFunctionDefinition(st)...)
	case *ast.ExpressionStatement:
		errors = append(errors, verifyExpressionStatement(st)...)
	case *ast.ReturnStatement:
		errors = append(errors, verifyReturnStatement(st)...)
	default:
		errors = append(errors, fmt.Errorf("unknown statement type: %T", stmt))
	}
}

func verifyVarStatement(vs *ast.VarStatement) []error {
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

func verifyFunctionDefinition(f *ast.FunctionDefinition) []error { e := []error{}; return e }
func verifyBlockStatement(f *ast.BlockStatement) []error         { e := []error{}; return e }
func verifyEntrypointFunctionDefinition(f *ast.EntrypointFunctionDefinition) []error {
	e := []error{}
	// check that this is the first definition of this identifier
	name := f.Name.Value
	if !contains(validEntryPointNames, name) {
		e = append(e, fmt.Errorf("unknown entrypoint name: %s", name))
	}
	entry := symtab.Entries[name]
	if entry != nil {
		e = append(e, fmt.Errorf("already defined %s", name))
	} else {
		symtab.Entries[name] = SymTabEntrypointFunction{}
	}
	return e
}
func verifyExpressionStatement(f *ast.ExpressionStatement) []error { e := []error{}; return e }
func verifyReturnStatement(f *ast.ReturnStatement) []error         { e := []error{}; return e }

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

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

var validEntryPointNames = []string{
	"main",
	// "init", ADD LATER!!!
	// "initOnce", ADD LATER!!!
}

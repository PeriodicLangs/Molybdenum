package verify

import (
	"fmt"

	"github.com/westsi/molybdenum/ast"
	"github.com/westsi/molybdenum/lex"
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
		switch stmt.NType() {
		case "VarStatement":
			errors = append(errors, VerifyVarStatement(stmt.(*ast.VarStatement))...)
		case "FunctionDefinition":
			errors = append(errors, VerifyFunctionDefinition(stmt.(*ast.FunctionDefinition))...)
		case "BlockStatement":
			errors = append(errors, VerifyBlockStatement(stmt.(*ast.BlockStatement))...)
		case "EntrypointFunctionDefinition":
			errors = append(errors, VerifyEntrypointFunctionDefinition(stmt.(*ast.EntrypointFunctionDefinition))...)
		case "ExpressionStatement":
			errors = append(errors, VerifyExpressionStatement(stmt.(*ast.ExpressionStatement))...)
		case "ReturnStatement":
			errors = append(errors, VerifyReturnStatement(stmt.(*ast.ReturnStatement))...)
		default:
			errors = append(errors, fmt.Errorf("unknown statement type: %s", stmt.NType()))
		}
	}

	return errors
}

func VerifyVarStatement(vs *ast.VarStatement) []error {
	e := []error{}
	name := vs.Name.Value
	entry := symtab.Entries[name]
	// check that this is the first definition of this identifier
	if entry != nil {
		e = append(e, fmt.Errorf("already defined %s", name))
	} else {
		symtab.Entries[name] = &SymTabVariable{Type: vs.Type.Value}
	}
	// check that the value assigned to it is of the correct type

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

func getLiteralType(lit lex.Token) string {
	switch lit {
	case lex.INTLITERAL:
		return "int"
	case lex.STRINGLITERAL:
		return "string"
	case lex.TRUE:
		return "bool"
	case lex.FALSE:
		return "bool"
	default:
		return "unknown"
	}
}

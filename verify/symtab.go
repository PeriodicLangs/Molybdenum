package verify

import (
	"fmt"
	"strings"
)

type SymTab struct {
	Entries map[string]SymTabEntry
}

type SymTabEntry interface {
	Literal() string
	String() string
}

type SymTabVariable struct {
	Type string
}

func (t SymTabVariable) Literal() string {
	return t.Type
}
func (t SymTabVariable) String() string {
	return t.Type
}

type SymTabFunction struct {
	ReturnType     string
	ParameterTypes []string
}

func (t SymTabFunction) Literal() string {
	return fmt.Sprintf("{%s, %s}", t.ReturnType, strings.Join(t.ParameterTypes, ", "))
}
func (t SymTabFunction) String() string {
	return fmt.Sprintf("{%s, %s}", t.ReturnType, strings.Join(t.ParameterTypes, ", "))
}

type SymTabEntrypointFunction struct{}

func (t SymTabEntrypointFunction) Literal() string { return "EPF" }

func (t SymTabEntrypointFunction) String() string { return "EPF" }

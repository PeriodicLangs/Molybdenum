package main

import "fmt"

func EvalFormatString(s string, v ...any) string {
	return fmt.Sprintf(s, v...)
}

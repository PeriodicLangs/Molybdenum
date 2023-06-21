package main

import (
	"fmt"
	"strconv"
)

func EvalFormatString(s string, v ...any) string {
	return fmt.Sprintf(s, v...)
}

func EvalMathExpr(exprs ...any) int {
	res := 0
	lastOp := ""
	for _, expr := range exprs {
		e := fmt.Sprint(expr)
		ii := isInt(e)
		if ii {
			ie, _ := strconv.Atoi(e)
			if lastOp == "" {
				res = ie
			} else {
				switch lastOp {
				case "+":
					res += ie
				case "-":
					res -= ie
				case "*":
					res *= ie
				case "/":
					res /= ie
				default:
					fmt.Println("unknown op")
				}
			}
		} else {
			lastOp = e
		}

	}
	return res
}

func isInt(s string) bool {
	_, e := strconv.Atoi(s)
	return e == nil
}

package leecode

import (
	"strconv"
)

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		switch token {
		case "+":
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = v1 + v2
		case "-":
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = v1 - v2
		case "*":
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = v1 * v2
		case "/":
			v1, v2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack[len(stack)-1] = v1 / v2
		default:
			val, _ := strconv.Atoi(token)
			stack = append(stack, val)
		}
	}
	return stack[0]
}

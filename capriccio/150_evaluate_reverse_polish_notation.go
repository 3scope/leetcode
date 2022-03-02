package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]string, 0, 8)
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "+" || tokens[i] == "-" || tokens[i] == "*" || tokens[i] == "/" {
			right := stack[len(stack)-1]
			left := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			stack = append(stack, oprationRPN(left, right, tokens[i]))
		} else {
			stack = append(stack, tokens[i])
		}
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}

func oprationRPN(x, y, oprator string) string {
	result := 0
	left, _ := strconv.Atoi(x)
	right, _ := strconv.Atoi(y)
	if oprator == "+" {
		result = left + right
	} else if oprator == "-" {
		result = left - right
	} else if oprator == "*" {
		result = left * right
	} else {
		result = left / right
	}
	return strconv.Itoa(result)
}

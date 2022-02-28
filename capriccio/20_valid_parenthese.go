package main

func isValid(s string) bool {
	stack := make([]byte, 0, 8)
	length := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, ')')
			length++
		} else if s[i] == ')' {
			if checkParentheses(length, stack, s[i]) {
				stack = stack[:length-1]
				length--
			} else {
				return false
			}
		} else if s[i] == '[' {
			stack = append(stack, ']')
			length++
		} else if s[i] == ']' {
			if checkParentheses(length, stack, s[i]) {
				stack = stack[:length-1]
				length--
			} else {
				return false
			}
		} else if s[i] == '{' {
			stack = append(stack, '}')
			length++
		} else if s[i] == '}' {
			if checkParentheses(length, stack, s[i]) {
				stack = stack[:length-1]
				length--
			} else {
				return false
			}
		}
	}
	return length == 0
}

func checkParentheses(length int, stack []byte, current byte) bool {
	if length <= 0 {
		return false
	}
	if stack[length-1] != current {
		return false
	}
	return true
}

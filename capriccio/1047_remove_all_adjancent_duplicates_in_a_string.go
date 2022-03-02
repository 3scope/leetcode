package main

func removeDuplicates(s string) string {
	stack := make([]byte, 0, 8)
	for i := 0; i < len(s); i++ {
		if len(stack) == 0 || stack[len(stack)-1] != s[i] {
			stack = append(stack, s[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

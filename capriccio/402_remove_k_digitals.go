package main

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	stack := make([]byte, 0)
	for i := 0; i < len(num); i++ {
		if len(stack) > 0 && num[i] >= stack[len(stack)-1] {
			stack = append(stack, num[i])
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > num[i] && k > 0 {
				k--
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, num[i])
		}
	}

	// Satisfy that at least k numbers are deleted.
	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	// To trim prefix.
	size := len(stack)
	for i := 0; i < size-1 && stack[0] == '0'; i++ {
		stack = stack[1:]
	}

	return string(stack)
}

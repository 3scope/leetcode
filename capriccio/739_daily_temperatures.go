package main

func dailyTemperatures(temperatures []int) []int {
	// The default value is zero.
	result := make([]int, len(temperatures))
	// Push the element with index 0 into the stack.
	// To store the index.
	stack := make([]int, 1)
	for i := 1; i < len(temperatures); i++ {
		if temperatures[i] < temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else if temperatures[i] == temperatures[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
				popIndex := stack[len(stack)-1]
				result[popIndex] = i - popIndex
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}

	return result
}

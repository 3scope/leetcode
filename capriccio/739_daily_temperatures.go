package main

func dailyTemperatures(temperatures []int) []int {
	// 默认值为0。
	result := make([]int, len(temperatures))
	// 需要求下一个更高的气温，至少等待的天数，因此需要存储数组的下标。
	stack := make([]int, 1)
	stack[0] = 0
	for i := 1; i < len(temperatures); i++ {
		// “stack[len(stack)-1]”是栈顶的元素，由于存放的是数组下标，需要将栈顶元素作为下标取到对应的值。
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

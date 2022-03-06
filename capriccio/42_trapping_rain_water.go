package main

func trap(height []int) int {
	// Special case.
	if len(height) <= 2 {
		return 0
	}
	// Store the index of array.
	stack := make([]int, 1)
	stack[0] = 0
	result := 0
	for i := 1; i < len(height); i++ {
		if height[stack[len(stack)-1]] == height[i] {
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else if height[stack[len(stack)-1]] > height[i] {
			stack = append(stack, i)
		} else {
			for len(stack) >= 2 {
				floor := stack[len(stack)-1]
				if height[floor] > height[i] {
					break
				}
				stack = stack[:len(stack)-1]
				leftHeight := stack[len(stack)-1]
				h := func() int {
					if height[leftHeight] > height[i] {
						return height[i]
					} else {
						return height[leftHeight]
					}
				}() - height[floor]
				w := i - leftHeight - 1
				result += h * w
			}
			for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, i)
		}
	}

	return result
}

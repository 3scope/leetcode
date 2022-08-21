package main

func trap(height []int) int {
	// 当只有两个柱子的时候，接不到任何雨水。
	if len(height) <= 2 {
		return 0
	}
	// 存放柱子的下标，通过下标可以计算宽度。
	stack := make([]int, 1)
	stack[0] = 0
	result := 0
	for i := 1; i < len(height); i++ {
		if height[stack[len(stack)-1]] == height[i] {
			// 弹出相同的元素，目的是为了防止重复计算。
			stack = stack[:len(stack)-1]
			stack = append(stack, i)
		} else if height[stack[len(stack)-1]] > height[i] {
			stack = append(stack, i)
		} else {
			// 当满足“高低高”的情况，就可以接到雨水。
			for len(stack) >= 2 {
				floor := stack[len(stack)-1]
				if height[floor] > height[i] {
					// 如果此时单调栈中没有比当前柱子矮的，那么就跳出循环。
					break
				}
				stack = stack[:len(stack)-1]
				leftHeight := stack[len(stack)-1]
				// “h”是接到的雨水的高度。
				h := func() int {
					if height[leftHeight] > height[i] {
						return height[i]
					} else {
						return height[leftHeight]
					}
				}() - height[floor]
				// “w”是接到的雨水的宽度。
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

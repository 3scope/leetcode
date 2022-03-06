package main

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	} else if len(heights) == 1 {
		return heights[0]
	}
	// Convert all cases to "big small big".
	temp := []int{0}
	heights = append(temp, heights...)
	heights = append(heights, 0)

	result := 0
	stack := make([]int, 1)
	stack[0] = 0
	for i := 1; i < len(heights); i++ {
		if heights[i] == heights[stack[len(stack)-1]] {
			stack[len(stack)-1]++
		} else if heights[i] > heights[stack[len(stack)-1]] {
			stack = append(stack, i)
		} else {
			for len(stack) > 2 {
				middle := stack[len(stack)-1]
				if heights[i] > heights[middle] {
					break
				}
				stack = stack[:len(stack)-1]
				left := stack[len(stack)-1]
				w := i - left - 1
				h := heights[middle]
				result = func() int {
					if result > w*h {
						return result
					} else {
						return w * h
					}
				}()
			}
		}
	}
}

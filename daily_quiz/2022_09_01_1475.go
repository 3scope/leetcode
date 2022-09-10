package main

// 使用单调栈维护候选答案的集合。
func finalPrices(prices []int) []int {
	// 存放结果。
	result := make([]int, len(prices))
	// 折扣是右边第一个比当前元素小的，因此单调栈从栈底到栈顶是单调递增的。
	// 也就是说从后往前处理“prices”数组，如果遇到了“j < k”并且“prices[j] < prices[k]”，那么“prices[j]”才是候选答案。
	stack := make([]int, 1, len(prices))
	// 初始值是为了避免单调栈出现没有元素的情况。
	stack[0] = 0
	// 从后往前遍历。
	for i := len(prices) - 1; i >= 0; i-- {
		cur := stack[len(stack)-1]
		if prices[i] >= cur {
			// 如果当前元素大于栈顶元素，那么折扣就是栈顶元素。
			result[i] = prices[i] - cur
			stack = append(stack, prices[i])
		} else {
			for len(stack) > 1 && prices[i] < cur {
				// 弹出栈顶元素。
				stack = stack[:len(stack)-1]
				cur = stack[len(stack)-1]
			}
			// 此时折扣就是栈顶元素。
			result[i] = prices[i] - cur
			// 之后再将元素放入栈中。
			stack = append(stack, prices[i])
		}
	}

	return result
}

package main

func maxProfit(prices []int) int {
	if len(prices) == 1 {
		return 0
	}

	stack := make([]int, 1)
	stack[0] = prices[0]
	result := 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(prices); i++ {
		if prices[i] < stack[len(stack)-1] {
			stack = append(stack, prices[i])
		} else if prices[i] == stack[len(stack)-1] {
			continue
		} else {
			result = max(result, prices[i]-stack[len(stack)-1])
		}
	}

	return result
}

func maxProfitDP(prices []int) int {
	if len(prices) == 1 {
		return 0
	}
	// The first value of dp[i] is the profit when holding the stock, and the second one is no longer holding.
	dp := make([][2]int, len(prices))
	// When buying the stock, and holding, the profit is negative.
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], -prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i])
	}

	return dp[len(dp)-1][1]
}

package main

func maxProfitWithCooldown(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	dp := make([][4]int, len(prices))
	// Status 0: Holding stock.
	dp[0][0] = -prices[0]
	// Status 1: Donnot have any stock.
	// - Had sold before.
	dp[0][1] = 0
	// - Selling now.
	dp[0][2] = 0
	// Status 2: Cooldown.
	dp[0][3] = 0

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = max(dp[i-1][0], max(dp[i-1][1]-prices[i], dp[i-1][3]-prices[i]))
		dp[i][1] = max(dp[i-1][1], dp[i-1][3])
		dp[i][2] = dp[i-1][0] + prices[i]
		dp[i][3] = dp[i-1][2]
	}

	return max(dp[len(dp)-1][1], max(dp[len(dp)-1][2], dp[len(dp)-1][3]))
}

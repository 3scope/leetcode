package main

func maxProfitThree(prices []int) int {
	// 0: No operations. This status also can delete or add to 121.
	// 1: First time to buy stock.(Status)
	// 2: First time to sell stock.
	// 3: Second time to buy stock.
	// 4: Second time to sell stock.
	dp := make([][5]int, len(prices))
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	dp[0][2] = 0
	dp[0][3] = -prices[0]
	dp[0][4] = 0

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		// First time to buy.
		dp[i][1] = max(dp[i-1][1], 0-prices[i])
		// First time to sell.
		dp[i][2] = max(dp[i-1][2], dp[i-1][1]+prices[i])
		// Second time.
		dp[i][3] = max(dp[i-1][3], dp[i-1][2]-prices[i])
		dp[i][4] = max(dp[i-1][4], dp[i-1][3]+prices[i])
	}

	return dp[len(dp)-1][4]
}

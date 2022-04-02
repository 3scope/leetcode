package main

func maxProfitFour(k int, prices []int) int {
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2*k+1)
	}
	// The time to buy the stock.
	for i := 1; i < 2*k+1; i += 2 {
		dp[0][i] = -prices[0]
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(prices); i++ {
		dp[i][0] = dp[i-1][0]
		for j := 1; j < 2*k+1; j += 2 {
			dp[i][j] = max(dp[i-1][j], dp[i-1][j-1]-prices[i])
			dp[i][j+1] = max(dp[i-1][j+1], dp[i-1][j]+prices[i])
		}
	}

	return dp[len(dp)-1][2*k]
}

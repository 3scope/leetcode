package main

func maxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 1 {
		return 0
	}
	result := 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i] >= minPrice && prices[i] <= minPrice+fee {
			continue
		} else {
			result += prices[i] - fee - minPrice
			// The profit of the next round is the difference between the current price and the next one.
			minPrice = prices[i] - fee
		}
	}

	return result
}

func maxProfitWithFeeDP(prices []int, fee int) int {
	if len(prices) <= 1 {
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
		dp[i][0] = max(dp[i-1][0], dp[i-1][1]-prices[i])
		dp[i][1] = max(dp[i-1][1], dp[i-1][0]+prices[i]-fee)
	}

	return dp[len(dp)-1][1]
}

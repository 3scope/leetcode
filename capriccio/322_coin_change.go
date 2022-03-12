package main

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 0
	// The goal is to find the minimum of the total, therefore, initialization needs to make the value of the dp array as large as possible.
	for i := 1; i <= amount; i++ {
		dp[i] = 10001
	}
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}
	if dp[amount] == 10001 {
		return -1
	}
	return dp[amount]
}

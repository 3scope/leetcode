package main

func numSquares(n int) int {
	dp := make([]int, n+1)
	// The iteration needs to be started with this initial value.
	dp[0] = 0
	for i := 1; i <= n; i++ {
		dp[i] = 10000
	}
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i*i <= n; i++ {
		for j := i * i; j <= n; j++ {
			dp[j] = min(dp[j], dp[j-i*i]+1)
		}
	}
	return dp[n]
}

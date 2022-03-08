package main

func integerBreak(n int) int {
	dp := make([]int, n+1)
	dp[2] = 1
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	// Because the value at index of 0 and 1 is nil, dp[i-j] at least is dp[2].
	for i := 3; i <= n; i++ {
		for j := 1; j < i-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

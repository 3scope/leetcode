package main

// The essence is an arrangement problem.
func climbStairs(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for j := 1; j <= n; j++ {
		for i := 1; i <= 2; i++ {
			if j-i >= 0 {
				dp[j] += dp[j-i]
			}
		}
	}
	return dp[n]
}

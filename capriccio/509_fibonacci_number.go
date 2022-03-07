package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	slow := 0
	fast := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = slow + fast
		slow = fast
		fast = result
	}

	return result
}

func fibArray(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

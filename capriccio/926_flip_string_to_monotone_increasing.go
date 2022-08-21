package main

func minFlipsMonoIncr(s string) int {
	// “dp[i][0]”代表将第“i”位最终置为“0”时，最小的变化次数。
	dp := make([][2]int, len(s))
	// 初始化“dp[i]”的值。
	if s[0] == '0' {
		dp[0][0] = 0
		dp[0][1] = 1
	} else {
		dp[0][0] = 1
		dp[0][1] = 0
	}
	min := func(x, y int) int {
		if x < y {
			return x
		} else {
			return y
		}
	}

	for i := 1; i < len(s); i++ {
		// 当前位是0。
		if s[i] == '0' {
			// 置位为0，那么变化次数不变，并且第“i-1”位必须是0。
			dp[i][0] = dp[i-1][0]
			// 置位为1，那么前面的既可以是0也可以是1。
			dp[i][1] = min(dp[i-1][0], dp[i-1][1]) + 1
		} else {
			// 当前位是1。
			// 置位为0，第“i-1”位必须是0。
			dp[i][0] = dp[i-1][0] + 1
			// 置位为1，那么前面的既可以是0也可以是1。
			dp[i][1] = min(dp[i-1][0], dp[i-1][1])
		}
	}

	return min(dp[len(s)-1][0], dp[len(s)-1][1])
}

package main

func isSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}
	if len(s) == len(t) && len(s) == 0 {
		return true
	}
	dp := make([][]int, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(t)+1)
	}

	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if s[i-1] == t[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// The string "t" remove a char, and then the length is "j-1".
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	return dp[len(s)][len(t)] == len(s)
}

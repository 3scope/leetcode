package main

func countSubstrings(s string) int {
	// Note that because of the definition of dp[i][j], j must be greater than or equal to i, then when filling dp[i][j], only the upper right half must be filled.
	dp := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	result := 0
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i <= 1 {
					dp[i][j] = true
					result++
				} else {
					if dp[i+1][j-1] {
						result++
					}
					dp[i][j] = dp[i+1][j-1]
				}
			} else {
				dp[i][j] = false
			}
		}
	}
	return result
}

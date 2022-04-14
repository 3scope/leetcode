package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(word2)+1)
	}

	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	maxSubsequenceLength := dp[len(word1)][len(word2)]
	return len(word1) - maxSubsequenceLength + len(word2) - maxSubsequenceLength
}

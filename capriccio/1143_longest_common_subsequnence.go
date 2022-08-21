package main

// 两个字符串的子序列都可以是原本不连续的。
func longestCommonSubsequence(text1 string, text2 string) int {
	// “dp[i][j]”从下标为0的字符开始，长度为“i”和长度为“j”的最长公共子序列长度的值。
	dp := make([][]int, len(text1)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(text2)+1)
	}

	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	// 当“i”或者“j”为0的时候，公共子序列长度为0。
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[i]); j++ {
			// 如果第一字符串的下标为“i-1”和第二个字符串的下标“j-1”的字符相同，那么长度为“i”和长度为“j”的最长公共子序列长度的值加1。
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				// 如果不相同，那么任意删除某一个字符串对应的字符，取最大值。
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(text1)][len(text2)]
}

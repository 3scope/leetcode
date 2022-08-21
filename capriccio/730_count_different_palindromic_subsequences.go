package main

func countPalindromicSubsequences(s string) int {
	// 考虑使用区间DP。
	// “dp[i][j]”代表下标从“i”到“j”闭区间子串包含的回文字符串的个数。
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		// 长度为1的子串都是回文字符串。
		dp[i][i] = 1
	}
	// 从长度为2的子串开始，计算回文字符串的个数。
	for length := 2; length <= len(s); length++ {
		// “i”和“j”为“dp”数组下标。
		for i := 0; i <= len(s)-length; i++ {
			j := i + length - 1
			// “i”和“j”所指的字符相同时。
			if s[i] == s[j] {
				// “dp[i][j]”的值由“dp[i+1][j-1]”推出。
				left := i + 1
				right := j - 1
				// 寻找下标从“i”到“j”闭区间子串中第一个和“s[i]”相同的字符。
				for left <= right {
					if s[left] == s[i] {
						break
					}
					left++
				}
				// 寻找下标从“j”到“i”闭区间子串中第一个和“s[i]”相同的字符。
				for left <= right {
					if s[right] == s[i] {
						break
					}
					right--
				}
				// 没有重复字符。
				if left > right {
					// 如果长度为2的话，即“j = i + 1”的话，那么“dp[i+1][j-1] = 0” => “dp[i][j] = 2”。
					dp[i][j] = 2*dp[i+1][j-1] + 2
				} else if left == right {
					// 有一个重复字符。
					dp[i][j] = 2*dp[i+1][j-1] + 1
				} else {
					dp[i][j] = 2*dp[i+1][j-1] - dp[left+1][right-1]
				}
			} else {
				// 如果“i”和“j”所指的字符不同时。
				dp[i][j] = dp[i][j-1] + dp[i+1][j] - dp[i+1][j-1]
			}
			// 对负数情况的处理。
			if dp[i][j] < 0 {
				dp[i][j] += 1e+9 + 7
			} else {
				dp[i][j] = dp[i][j] % (1e+9 + 7)
			}
		}
	}
	return dp[0][len(s)-1]
}

package main

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1)
	dp[0] = true
	// 背包为目标串，容量是 len(s)，物品为单词，物品重量为 len(wordDict[i])。
	// 只有当截取子串前的字符串可以被表示，并且自身也可以被表示的时候，整个长度为 j 的字符串才能够被表示出来。
	for j := 1; j <= len(s); j++ {
		for i := 0; i < len(wordDict); i++ {
			if j >= len(wordDict[i]) && dp[j-len(wordDict[i])] && s[j-len(wordDict[i]):j] == wordDict[i] {
				dp[j] = true
			}
		}
	}

	return dp[len(s)]
}

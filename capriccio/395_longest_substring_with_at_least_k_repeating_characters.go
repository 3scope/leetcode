package main

import "strings"

// 判断是否能用滑动窗口，主要看是不是可以分析出满足条件或者不满足条件时，窗口如何滑动。
func longestSubstring(s string, k int) int {
	// 终止条件。
	if len(s) < k {
		return 0
	}

	// 单层逻辑处理。
	byteCount := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		byteCount[s[i]]++
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}

	for b, v := range byteCount {
		// 出现了字符数量不足k的，当前字符串不满足要求。
		if v < k {
			// 会出现空字符串，递归返回0。
			temp := strings.Split(s, string(b))
			result := 0
			for i := 0; i < len(temp); i++ {
				result = max(result, longestSubstring(temp[i], k))
			}
			return result
		}
	}
	// 如果所有的字符数量都高于k，那么就直接返回s的长度。
	return len(s)
}

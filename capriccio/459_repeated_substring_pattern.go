package main

func repeatedSubstringPattern(s string) bool {
	// If the string length is one, the substring of it is null.
	if len(s) == 1 || len(s) == 0 {
		return false
	}
	next := GetPrefixNext(s)
	// If the last digit of the next is zero, then prove that the string is not repeated.
	if next[len(s)-1] != 0 && len(s)%(len(s)-next[len(s)-1]) == 0 {
		return true
	} else {
		return false
	}
}

func KMP(haystack string, pattern string) int {
	if len(pattern) == 0 {
		return 0
	}
	next := GetPrefixNext(pattern)
	var i, j int
	for i < len(haystack) {
		if haystack[i] == pattern[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = next[j-1]
		}
		// Match successfully.
		if j == len(pattern) {
			return i - j
		}
	}
	return -1
}

func GetPrefixNext(pattern string) []int {
	prefixNext := make([]int, len(pattern))
	// Next 数组即是前缀表（最长相等前后缀表，前缀和后缀都不包括自身），存储着子串最长相等前后缀的长度。
	prefixNext[0] = 0
	// 求 Next 数组的过程，本质上是对自身进行模式匹配。
	// 使用变量 i 遍历整个模式串，同时 i 还是子串*后缀*的最后一个字符=>  abc ab
	// 														   	    ab cab
	// 此时 len(pattern) = 5, i = 4，最长相等前后缀是 ab，后缀的最后一个字符即为 pattern[i] = b。
	// 变量 j 是*前缀*的最后一个字符，因此可以代表*最长相等前后缀长度*，它和 KMP 算法中的 j 的作用是一致的，是模式串中的指针，此时 j = 1, patttern[j] = b
	i := 1
	j := 0
	for i < len(pattern) {
		if pattern[i] == pattern[j] {
			prefixNext[i] = j + 1
			i++
			j++
		} else if j == 0 {
			prefixNext[i] = 0
			i++
		} else {
			j = prefixNext[j-1]
		}
	}

	return prefixNext
}

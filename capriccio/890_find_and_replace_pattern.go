package main

// 建立从“pattern”到“words[i]”之间的映射。
func findAndReplacePattern(words []string, pattern string) []string {
	// 存放结果。
	result := make([]string, 0)
	for i := 0; i < len(words); i++ {
		// 每个单词都新建一个“map”，相当于一种清空操作。
		patternToWords := make(map[byte]byte)
		// 判断“words”中的某个字符是否已经映射过。
		isExist := make(map[byte]bool)
		// 第一个字符第一次出现一定在首位。
		patternToWords[pattern[0]] = words[i][0]
		isExist[words[i][0]] = true
		// 判断是否满足模式。
		flag := true
		for j := 1; j < len(pattern); j++ {
			// 如果对应的字符是第一次出现，那么记录下来。
			if _, ok := patternToWords[pattern[j]]; !ok {
				// 如果“words”中的字符已经映射过。
				if isExist[words[i][j]] {
					flag = false
					break
				}
				patternToWords[pattern[j]] = words[i][j]
				isExist[words[i][j]] = true
			} else {
				// 不是第一次出现，进行匹配。
				if patternToWords[pattern[j]] != words[i][j] {
					flag = false
					break
				}
			}
		}
		if flag {
			result = append(result, words[i])
		}
	}
	return result
}

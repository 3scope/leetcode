package main

func lengthOfLongestSubstring(s string) int {
	max, start, end := 0, 0, 0
	charSet := make(map[byte]int)

	for end < len(s) {
		if _, ok := charSet[s[end]]; ok {
			delete(charSet, s[start])
			start++
		} else {
			charSet[s[end]]++
			end++
			max = maxInt(max, len(charSet))
		}
	}

	return max
}

func maxInt(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}

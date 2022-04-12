package main

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right := 0, 0
	result := 1
	max := func(i, j int) int {
		if i > j {
			return i
		} else {
			return j
		}
	}
	isExist := make(map[byte]bool)
	for right < len(s) {
		if isExist[s[right]] {
			isExist[s[left]] = false
			left++
		} else {
			isExist[s[right]] = true
			right++
		}
		result = max(right-left, result)
	}

	return result
}

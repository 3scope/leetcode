package main

func numberOfLines(widths []int, s string) []int {
	length := 100
	lineCount := 1
	for i := 0; i < len(s); i++ {
		char := s[i]
		if length >= widths[char-'a'] {
			length -= widths[char-'a']
		} else {
			lineCount++
			length = 100 - widths[char-'a']
		}
	}
	return []int{lineCount, 100 - length}
}

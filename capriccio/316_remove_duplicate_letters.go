package main

func removeDuplicateLetters(s string) string {
	if len(s) == 1 {
		return s
	}
	// Save the count of each letter which appears in s.
	letters := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		letters[s[i]]++
	}

	stack := make([]byte, 1)
	// To judge whether letter is exist.
	isExist := make(map[byte]bool)
	stack[0] = s[0]
	isExist[s[0]] = true
	letters[s[0]]--
	for i := 1; i < len(s); i++ {
		letters[s[i]]--
		if isExist[s[i]] {
			continue
		}
		if s[i] > stack[len(stack)-1] {
			stack = append(stack, s[i])
			isExist[s[i]] = true
		} else if s[i] == stack[len(stack)-1] {
			continue
		} else {
			for len(stack) > 0 && s[i] < stack[len(stack)-1] {
				if letters[stack[len(stack)-1]] <= 0 {
					break
				}
				isExist[stack[len(stack)-1]] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, s[i])
			isExist[s[i]] = true
		}
	}

	return string(stack)
}

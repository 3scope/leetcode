package main

func partition(s string) [][]string {
	result := make([][]string, 0)
	temp := make([]string, 0)
	partitionBacktracking(s, 0, &result, &temp)

	return result
}

func partitionBacktracking(s string, startIndex int, result *[][]string, temp *[]string) {
	if startIndex == len(s) {
		new := make([]string, len(*temp))
		for i := 0; i < len(new); i++ {
			new[i] = (*temp)[i]
		}
		*result = append(*result, new)
	}

	for i := startIndex; i < len(s); i++ {
		if isPalindrome(string(s[startIndex : i+1])) {
			*temp = append(*temp, string(s[startIndex:i+1]))
			partitionBacktracking(s, i+1, result, temp)
			*temp = (*temp)[:len(*temp)-1]
		} else {
			continue
		}
	}
}

func isPalindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

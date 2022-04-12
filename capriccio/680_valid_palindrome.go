package main

func validPalindrome(s string) bool {
	if len(s) <= 2 {
		return true
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return isPalindromeString(s[i:j]) || isPalindromeString(s[i+1:j+1])
		}
	}
	return true
}

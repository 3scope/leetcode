package main

func isPalindrome(x int) bool {
	// Special cases.
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revered := 0
	for x > revered {
		revered = revered*10 + x%10
		x = x / 10
	}

	return x == revered || x == revered/10
}

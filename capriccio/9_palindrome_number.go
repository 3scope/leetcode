package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	} else if x%10 == 0 {
		return false
	}

	reverseNumber := 0
	for reverseNumber < x {
		reverseNumber = reverseNumber*10 + x%10
		x = x / 10
	}

	// Even lenght of x.
	if reverseNumber == x {
		return true
	}

	reverseNumber = reverseNumber / 10
	return reverseNumber == x
}

package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	letters := []string{
		"",
		"",
		"abc",
		"def",
		"ghi",
		"jkl",
		"mno",
		"pqrs",
		"tuv",
		"wxyz",
	}
	result := make([]string, 0)
	temp := make([]byte, 0)
	letterCombinationsBacktracking(digits, 0, letters, &result, &temp)

	return result
}

func letterCombinationsBacktracking(digits string, digitsIndex int, letters []string, result *[]string, temp *[]byte) {
	if len(*temp) == len(digits) {
		*result = append(*result, string(*temp))
		return
	}

	// The "digitsIndex" is to control recursion depth.
	theDigit := digits[digitsIndex] - '0'
	letterSet := letters[theDigit]
	for i := 0; i < len(letterSet); i++ {
		*temp = append(*temp, []byte(letterSet)[i])
		letterCombinationsBacktracking(digits, digitsIndex+1, letters, result, temp)
		*temp = (*temp)[:len(*temp)-1]
	}
}

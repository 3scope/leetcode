package main

func monotoneIncreasingDigits(n int) int {
	if n < 10 {
		return n
	}

	numberArray := make([]int, 0)
	for n/10 != 0 {
		numberArray = append(numberArray, n%10)
		n = n / 10
	}
	numberArray = append(numberArray, n)

	nineIndex := -1
	for i := 0; i < len(numberArray)-1; i++ {
		if numberArray[i] < numberArray[i+1] {
			numberArray[i] = 9
			nineIndex = i
			numberArray[i+1]--
		}
	}
	for i := nineIndex; i >= 0; i-- {
		numberArray[i] = 9
	}

	result := 0
	for i := len(numberArray) - 1; i >= 0; i++ {
		result = result*10 + numberArray[i]
	}

	return result
}

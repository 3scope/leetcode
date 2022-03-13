package main

func candy(ratings []int) int {
	if len(ratings) == 1 {
		return 1
	}
	resultSlice := make([]int, len(ratings))
	resultSlice[0] = 1
	resultSlice[len(ratings)-1] = 1
	for i := 1; i < len(ratings); i++ {
		resultSlice[i] = 1
		if ratings[i] > ratings[i-1] {
			resultSlice[i] = resultSlice[i-1] + 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if resultSlice[i+1]+1 > resultSlice[i] {
				resultSlice[i] = resultSlice[i+1] + 1
			}
		}
	}

	result := 0
	for i := 0; i < len(resultSlice); i++ {
		result += resultSlice[i]
	}
	return result
}

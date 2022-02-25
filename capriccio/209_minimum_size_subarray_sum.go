package main

func minSubArrayLen(target int, nums []int) int {
	result := 100001
	length := len(nums)
	i, j := 0, 0
	for j < length {
		sum := 0
		for index := i; index <= j; index++ {
			sum += nums[index]
		}
		if sum >= target {
			result = min(result, j-i+1)
			i++
		} else {
			j++
		}
	}
	if result == 100001 {
		return 0
	}
	return result
}

func min(x, y int) int {
	if x > y {
		return y
	} else {
		return x
	}
}

package main

func findLengthOfLCISGreedy(nums []int) int {
	result := 1
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp++
		} else {
			temp = 1
		}
		result = max(result, temp)
	}

	return result
}

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
		if result < dp[i] {
			result = dp[i]
		}
	}

	return result
}

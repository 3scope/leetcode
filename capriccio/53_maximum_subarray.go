package main

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := -10000
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}
	return result
}

func maxSubArrayDP(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	result := dp[0]
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 1; i < len(nums); i++ {
		dp[i] = max(nums[i], dp[i-1]+nums[i])
		result = max(result, dp[i])
	}

	return result
}

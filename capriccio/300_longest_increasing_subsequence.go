package main

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	result := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
			result = max(result, dp[i])
		}
	}

	return result
}

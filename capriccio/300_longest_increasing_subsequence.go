package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := make([]int, len(nums))
	dp[0] = 1

	return dp[len(dp)-1]
}

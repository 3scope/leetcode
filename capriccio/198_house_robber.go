package main

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	count := len(nums)
	dp := make([]int, count)
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < count; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[count-1]
}

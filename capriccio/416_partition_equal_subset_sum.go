package main

func canPartition(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%2 == 1 {
		return false
	}
	target := sum / 2
	// The capacity of backpack is target
	dp := make([]int, target+1)
	// dp[0] = 0
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	for i := 0; i < len(nums); i++ {
		// for j := target; j >= nums[i]; j--
		for j := target; j >= 0; j-- {
			if j-nums[i] < 0 {
				continue
			}
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
		}
	}
	return dp[target] == target
}

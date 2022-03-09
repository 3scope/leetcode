package main

func findTargetSumWays(nums []int, target int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if target > sum || target < -sum {
		// Have no solution.
		return 0
	}
	if (target+sum)%2 == 1 {
		return 0
	}

	t := (target + sum) / 2
	dp := make([]int, t+1)
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		for j := t; j >= nums[i]; j-- {
			dp[j] += dp[j-nums[i]]
		}
	}

	return dp[t]
}

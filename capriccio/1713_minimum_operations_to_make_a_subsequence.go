package main

// 需要填充“arr”数组，直到“target”可以成为“arr”的一个子序列。
func minOperations(target []int, arr []int) int {
	// 存放所有“target”中的值和其下标顺序。
	targetIndex := make(map[int]int)
	for i := 0; i < len(target); i++ {
		targetIndex[target[i]] = i
	}
	// “nums”用于存放“arr”中已有的“target”值的下标，可以代表“target”中的正确顺序。
	nums := make([]int, 0, len(targetIndex))
	for i := 0; i < len(arr); i++ {
		if value, ok := targetIndex[arr[i]]; ok {
			nums = append(nums, value)
		}
	}
	// 如果“nums”为空的话，那么直接返回“target”的长度。
	if len(nums) == 0 {
		return len(target)
	}

	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	// 得到“nums”中单调递增子序列的长度，可以求的“arr”中正确的相对顺序。
	dp := make([]int, len(nums))
	// 赋初始值。
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}
	result := 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// “result”中存放“dp”的最大值。
		result = max(result, dp[i])
	}
	return len(target) - result
}

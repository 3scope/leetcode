package main

func wiggleMaxLength(nums []int) int {
	if len(nums) == 1 {
		return 1
	}

	preDifference := nums[1] - nums[0]
	curDifference := 0
	result := 2
	if preDifference == 0 {
		result = 1
	}

	for i := 2; i < len(nums); i++ {
		curDifference = nums[i] - nums[i-1]
		// Only initial value of "preDifference" can be 0.
		if (curDifference > 0 && preDifference <= 0) || (curDifference < 0 && preDifference >= 0) {
			result++
			preDifference = curDifference
		}
	}

	return result
}

func wiggleMaxLengthDP(nums []int) int {
	if len(nums) == 1 {
		return 1
	}
	dp := make([][2]int, len(nums))
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	// The first value of dp[i] is the length of wiggle subsequence when i as a regional minimum.
	// The second value of dp[i] is the length of wiggle subsequence when i as a regional maximum.
	dp[0][0] = 1
	dp[0][1] = 1
	for i := 1; i < len(nums); i++ {

		for j := 0; j < i; j++ {
			if nums[i]-nums[j] < 0 {
				dp[i][0] = max(dp[i][0], dp[j][1]+1)
			} else if nums[i]-nums[j] == 0 {
				dp[i][0] = max(dp[j][0], dp[i][0])
				dp[i][1] = max(dp[j][1], dp[i][1])
			} else {
				dp[i][1] = max(dp[i][1], dp[j][0]+1)
			}
		}
	}
	return max(dp[len(dp)-1][0], dp[len(dp)-1][1])
}

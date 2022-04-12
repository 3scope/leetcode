package main

func findLength(nums1 []int, nums2 []int) int {
	dp := make([][]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		dp[i] = make([]int, len(nums2))
	}
	// Init
	result := 0
	for j := 0; j < len(nums2); j++ {
		if nums1[0] == nums2[j] {
			dp[0][j] = 1
			result = 1
		}
	}
	for i := 1; i < len(nums1); i++ {
		if nums1[i] == nums2[0] {
			dp[i][0] = 1
			result = 1
		}
	}

	for i := 1; i < len(nums1); i++ {
		for j := 1; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				dp[i][j] = dp[i-1][j-1] + 1
			}
			if result < dp[i][j] {
				result = dp[i][j]
			}
		}
	}

	return result
}

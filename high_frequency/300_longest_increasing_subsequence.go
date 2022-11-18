package main

// 递增的子序列可以是不连续的，因此对于任意一位字符，前驱可以考虑在它之前的所有字符。
func lengthOfLIS(nums []int) int {
	// “dp[i]”代表以下标“i”结尾的字符，它的最长递增子序列为“dp[i]”。
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
	// “result”最小值为1。
	result := 1
	for i := 1; i < len(nums); i++ {
		// 遍历之前出现的所有字符。
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		// “result”保存其中的最大值。
		result = max(result, dp[i])
	}

	return result
}

// 遍历之前出现的字符的时候，可以使用二分查找降低时间复杂度。
func lengthOfLISBinarySearch(nums []int) int {
	// “tails”数组用于保存前一种状态可能出现的最长递增子序列。
	// “tails[i]”的值代表长度为“i+1”的子序列尾部元素的值。
	tails := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		// 如果“tails”数组的最后一个元素要比当前元素小，那么“tails”数组的长度加一。
		if i == 0 || nums[i] > tails[len(tails)-1] {
			tails = append(tails, nums[i])
			continue
		}
		// 否则使用二分查找，定位元素放置的位置。
		left, right := 0, len(tails)-1
		for left <= right {
			middle := left + (right-left)/2
			if tails[middle] > nums[i] {
				right = middle - 1
			} else if tails[middle] < nums[i] {
				left = middle + 1
			} else {
				// 处理情况与“tails[middle] > nums[i]”相同。
				right = middle - 1
			}
		}
		tails[left] = nums[i]
	}
	return len(tails)
}

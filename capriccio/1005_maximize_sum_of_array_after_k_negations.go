package main

import (
	"sort"
)

func largestSumAfterKNegations(nums []int, k int) int {
	// Sort by absolute value.
	sort.Slice(nums, func(i, j int) bool {
		if nums[i] < 0 {
			i = -nums[i]
		} else {
			i = nums[i]
		}
		if nums[j] < 0 {
			j = -nums[j]
		} else {
			j = nums[j]
		}
		return i < j
	})
	for i := len(nums) - 1; i >= 0; i-- {
		if k <= 0 {
			break
		}
		if nums[i] < 0 {
			nums[i] = -nums[i]
			k--
		}
	}
	if k > 0 {
		k = k % 2
	}
	for k != 0 {
		nums[0] = -nums[0]
		k--
	}

	result := 0
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}

	return result
}

package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	sort.Ints(nums)
	index := len(nums) - 1
	if nums[len(nums)-1] > target {
		return false
	}
	if nums[len(nums)-1] == target {
		k--
		index--
	}
	subset := make([]int, k)
	return partitionDFS(nums, index, subset, target)
}

func partitionDFS(nums []int, index int, subset []int, target int) bool {
	if index < 0 {
		return true
	}
	// In this turn, "nums[index]" is selected.
	selected := nums[index]
	// The value of "len(subset) is k."
	for i := 0; i < len(subset); i++ {
		if subset[i]+selected <= target {
			subset[i] += selected
			if partitionDFS(nums, index-1, subset, target) {
				return true
			}
			// Otherwise, backtracking is required(else).
			subset[i] -= selected
		}
	}
	return false
}

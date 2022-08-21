package main

import "sort"

func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) < k {
		return false
	}
	sum := 0
	// 求总和，总和如果不能整除“k”，那么就不能分成“k”份相等的。
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if sum%k != 0 {
		return false
	}
	target := sum / k
	// 求组合，可以先排序，之后可以去重剪枝。
	sort.Ints(nums)
	index := len(nums) - 1
	if nums[len(nums)-1] > target {
		return false
	}
	// 如果有单独的一个元素等于“target”的话，把该元素单独拎出去。
	if nums[len(nums)-1] == target {
		k--
		index--
	}
	subset := make([]int, k)
	return partitionDFS(nums, index, subset, target)
}

// 递归树的深度为“len(nums)”，每个节点的宽度为“k”。
func partitionDFS(nums []int, index int, subset []int, target int) bool {
	if index < 0 {
		return true
	}
	// 当前选择的元素是下标为“index”的。
	selected := nums[index]
	for i := 0; i < len(subset); i++ {
		if subset[i]+selected <= target {
			subset[i] += selected
			// 递归。
			if partitionDFS(nums, index-1, subset, target) {
				return true
			}
			// 回溯。
			subset[i] -= selected
		}
	}
	return false
}

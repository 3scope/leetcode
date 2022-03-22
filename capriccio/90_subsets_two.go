package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	sort.Ints(nums)
	// Subset deduplication must be sorted.
	subsetsWithDupBacktrackingTwo(nums, 0, &result, &temp)

	return result
}

func subsetsWithDupBacktracking(nums []int, startIndex int, result *[][]int, temp *[]int) {
	new := make([]int, len(*temp))
	for i := 0; i < len(*temp); i++ {
		new[i] = (*temp)[i]
	}
	*result = append(*result, new)

	if startIndex >= len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		if i > startIndex && nums[i] == nums[i-1] {
			continue
		} else {
			*temp = append(*temp, nums[i])
			subsetsWithDupBacktracking(nums, i+1, result, temp)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}

func subsetsWithDupBacktrackingTwo(nums []int, startIndex int, result *[][]int, temp *[]int) {
	new := make([]int, len(*temp))
	copy(new, *temp)
	*result = append(*result, new)

	if startIndex >= len(nums) {
		return
	}

	// Stores whether duplicate nodes are used in this layer.
	levelUsed := make(map[int]bool)
	for i := startIndex; i < len(nums); i++ {
		if levelUsed[nums[i]] {
			continue
		} else {
			levelUsed[nums[i]] = true
			*temp = append(*temp, nums[i])
			subsetsWithDupBacktrackingTwo(nums, i+1, result, temp)
			*temp = (*temp)[:len(*temp)-1]
		}
	}
}

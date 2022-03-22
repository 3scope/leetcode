package main

import "sort"

func permuteUnique(nums []int) [][]int {
	result := make([][]int, 0)
	subResult := make([]int, 0)
	usedIndex := make(map[int]bool)
	sort.Ints(nums)
	permuteUniqueBacktracking(nums, &usedIndex, &result, &subResult)

	return result
}

func permuteUniqueBacktracking(nums []int, usedIndex *map[int]bool, result *[][]int, subResult *[]int) {
	if len(*subResult) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, *subResult)
		*result = append(*result, temp)
		return
	}

	levelUsed := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if (*usedIndex)[i] || levelUsed[nums[i]] {
			continue
		}
		levelUsed[nums[i]] = true
		(*usedIndex)[i] = true
		*subResult = append(*subResult, nums[i])
		permuteUniqueBacktracking(nums, usedIndex, result, subResult)
		(*usedIndex)[i] = false
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}

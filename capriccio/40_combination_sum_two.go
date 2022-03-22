package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	// Need to sort silce.
	sort.Ints(candidates)
	combinationSum2Backtracking(candidates, 0, target, 0, &result, &temp)

	return result
}

func combinationSum2Backtracking(candidates []int, startIndex int, target int, sum int, result *[][]int, temp *[]int) {
	if sum == target {
		new := make([]int, len(*temp))
		for i := 0; i < len(*temp); i++ {
			new[i] = (*temp)[i]
		}
		*result = append(*result, new)
	}

	if sum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		if i > startIndex && candidates[i] == candidates[i-1] {
			continue
		}
		sum += candidates[i]
		*temp = append(*temp, candidates[i])
		combinationSum2Backtracking(candidates, i+1, target, sum, result, temp)
		sum -= candidates[i]
		*temp = (*temp)[:len(*temp)-1]
	}
}

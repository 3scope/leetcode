package main

func combinationSum(candidates []int, target int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	combinationSumBacktracking(candidates, target, 0, 0, &result, &temp)

	return result
}

func combinationSumBacktracking(candidates []int, target int, sum int, startIndex int, result *[][]int, temp *[]int) {
	if sum == target {
		new := make([]int, len(*temp))
		for i := 0; i < len(new); i++ {
			new[i] = (*temp)[i]
		}
		*result = append(*result, new)
		return
	}

	if sum > target {
		return
	}

	for i := startIndex; i < len(candidates); i++ {
		sum += candidates[i]
		*temp = append(*temp, candidates[i])
		combinationSumBacktracking(candidates, target, sum, i, result, temp)
		sum -= candidates[i]
		*temp = (*temp)[:len(*temp)-1]
	}
}

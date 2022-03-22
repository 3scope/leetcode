package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	subsetsBacktracking(nums, 0, &result, &temp)

	return result
}

func subsetsBacktracking(nums []int, startIndex int, result *[][]int, temp *[]int) {
	new := make([]int, len(*temp))
	for i := 0; i < len(*temp); i++ {
		new[i] = (*temp)[i]
	}
	*result = append(*result, new)
	if startIndex >= len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		*temp = append(*temp, nums[i])
		subsetsBacktracking(nums, i+1, result, temp)
		*temp = (*temp)[:len(*temp)-1]
	}
}

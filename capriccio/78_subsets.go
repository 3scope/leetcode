package main

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	temp := make([]int, 0)
	subsetsBacktracking(nums, 0, &result, &temp)

	return result
}

func subsetsBacktracking(nums []int, startIndex int, result *[][]int, subResult *[]int) {
	temp := make([]int, len(*subResult))
	copy(temp, *subResult)
	*result = append(*result, temp)
	if startIndex >= len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		*subResult = append(*subResult, nums[i])
		subsetsBacktracking(nums, i+1, result, subResult)
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}

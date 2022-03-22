package main

func permute(nums []int) [][]int {
	result := make([][]int, 0)
	subResult := make([]int, 0)
	used := make(map[int]bool)
	permuteBacktracking(nums, &used, &result, &subResult)

	return result
}

func permuteBacktracking(nums []int, used *map[int]bool, result *[][]int, subResult *[]int) {
	if len(*subResult) == len(nums) {
		temp := make([]int, len(nums))
		copy(temp, *subResult)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if (*used)[nums[i]] {
			continue
		}
		(*used)[nums[i]] = true
		*subResult = append(*subResult, nums[i])
		permuteBacktracking(nums, used, result, subResult)
		(*used)[nums[i]] = false
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}

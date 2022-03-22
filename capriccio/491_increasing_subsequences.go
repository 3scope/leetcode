package main

func findSubsequences(nums []int) [][]int {
	result := make([][]int, 0)
	subResult := make([]int, 0)
	findSubsequencesBacktracking(nums, 0, &result, &subResult)

	return result
}

func findSubsequencesBacktracking(nums []int, startIndex int, result *[][]int, subResult *[]int) {
	if len(*subResult) >= 2 {
		new := make([]int, len(*subResult))
		copy(new, *subResult)
		*result = append(*result, new)
	}

	if startIndex >= len(nums) {
		return
	}

	levelUsed := make(map[int]bool)
	for i := startIndex; i < len(nums); i++ {
		if levelUsed[nums[i]] || (len(*subResult) > 0 && nums[i] < (*subResult)[len(*subResult)-1]) {
			continue
		}
		levelUsed[nums[i]] = true
		*subResult = append(*subResult, nums[i])
		findSubsequencesBacktracking(nums, i+1, result, subResult)
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}

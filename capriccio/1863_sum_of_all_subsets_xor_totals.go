package main

func subsetXORSum(nums []int) int {
	result := 0
	subResult := make([]int, 0)
	findSubsets(nums, 0, &result, &subResult)

	return result
}

func findSubsets(nums []int, startIndex int, result *int, subResult *[]int) {
	if len(*subResult) > 0 {
		xorSum := (*subResult)[0]
		for i := 1; i < len(*subResult); i++ {
			xorSum = xorSum ^ (*subResult)[i]
		}
		*result += xorSum
	}
	if startIndex == len(nums) {
		return
	}

	for i := startIndex; i < len(nums); i++ {
		*subResult = append(*subResult, nums[i])
		findSubsets(nums, i+1, result, subResult)
		*subResult = (*subResult)[:len(*subResult)-1]
	}
}

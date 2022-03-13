package main

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	result := -10000
	count := 0
	for i := 0; i < len(nums); i++ {
		count += nums[i]
		if count > result {
			result = count
		}
		if count < 0 {
			count = 0
		}
	}
	return result
}

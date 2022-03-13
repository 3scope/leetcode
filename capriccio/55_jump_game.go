package main

func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	cover := nums[0]
	result := false
	for i := 1; i <= cover; i++ {
		if i+nums[i] > cover {
			cover = i + nums[i]
		}
		if cover >= len(nums) {
			result = true
			break
		}
	}
	return result
}

package main

func sortedSquares(nums []int) []int {
	// 使用双指针法，指向数组两端，比较两端绝对值的大小。
	left, right := 0, len(nums)-1
	// 求绝对值。
	abs := func(x int) int {
		if x < 0 {
			return -x
		} else {
			return x
		}
	}
	// 存储结果。
	result := make([]int, len(nums))
	index := len(result) - 1
	for left <= right {
		if abs(nums[left]) > abs(nums[right]) {
			result[index] = nums[left] * nums[left]
			left++
		} else {
			result[index] = nums[right] * nums[right]
			right--
		}
	}
	return result
}

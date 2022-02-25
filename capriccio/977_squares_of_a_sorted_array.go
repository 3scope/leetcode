package main

func sortedSquares(nums []int) []int {
	i, j := 0, len(nums)-1
	result := make([]int, len(nums))
	index := j
	for ; i <= j; index-- {
		if nums[i]*nums[i] > nums[j]*nums[j] {
			result[index] = nums[i] * nums[i]
			i++
		} else {
			result[index] = nums[j] * nums[j]
			j--
		}
	}
	return result
}

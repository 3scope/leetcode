package main

func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = -1
	}

	stack := make([]int, 1, 8)
	stack[0] = 0
	// Iterate over each element twice.
	for i := 0; i < 2*len(result); i++ {
		index := stack[len(stack)-1]
		if nums[i%len(nums)] < nums[index] {
			stack = append(stack, i%len(nums))
		} else if nums[i%len(nums)] == nums[index] {
			stack = append(stack, i%len(nums))
		} else {
			for len(stack) > 0 && nums[i%len(nums)] > nums[stack[len(stack)-1]] {
				index = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				result[index] = nums[i%len(nums)]
			}
			stack = append(stack, i%len(nums))
		}
	}

	return result
}

package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	getIndex := make(map[int]int)
	// To store the index of number in nums1.
	for i := 0; i < len(nums1); i++ {
		getIndex[nums1[i]] = i
	}

	// To set the default result number.
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = -1
	}

	stack := make([]int, 1, 8)
	stack[0] = nums2[0]
	// Iterate over each element of the array to find its next largest value.
	for i := 1; i < len(nums2); i++ {
		if nums2[i] < stack[len(stack)-1] {
			stack = append(stack, nums2[i])
		} else if nums2[i] == stack[len(stack)-1] {
			stack = append(stack, nums2[i])
		} else {
			for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
				ele := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if index, ok := getIndex[ele]; ok {
					result[index] = nums2[i]
				}
			}
			stack = append(stack, nums2[i])
		}
	}

	return result
}

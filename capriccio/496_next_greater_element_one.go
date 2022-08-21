package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	// 如果找不到下一个更大的元素，需要把结果设置为-1。
	subResult := make(map[int]int)
	for i := 0; i < len(nums1); i++ {
		subResult[nums1[i]] = -1
	}

	stack := make([]int, 1)
	stack[0] = nums2[0]
	// 遍历“nums2”，因为是在“nums2”里面找下一个更大元素。
	for i := 1; i < len(nums2); i++ {
		if nums2[i] < stack[len(stack)-1] {
			stack = append(stack, nums2[i])
		} else if nums2[i] == stack[len(stack)-1] {
			stack = append(stack, nums2[i])
		} else {
			for len(stack) > 0 && nums2[i] > stack[len(stack)-1] {
				ele := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				subResult[ele] = nums2[i]
			}
			stack = append(stack, nums2[i])
		}
	}

	// 将“map”里的值按照原顺序复制到“result”中。
	result := make([]int, len(nums1))
	for i := 0; i < len(nums1); i++ {
		result[i] = subResult[nums1[i]]
	}

	return result
}

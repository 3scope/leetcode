package main

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	result := make([]int, 0)
	inter := make(map[int]bool)
	// 存储相同的数值。
	// 先得到数组1中有哪些数。
	for _, value := range nums1 {
		inter[value] = true
	}
	// 遍历数组2，将存在的数置位。
	// 存在的数置位后，也可以反向去重。
	for _, value := range nums2 {
		if inter[value] {
			inter[value] = false
			result = append(result, value)
		}
	}

	return result
}

package main

// 使用一个“map”来存储第一个和第二个数组之和，“key”为两数之和，“value”为出现的次数。
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// 第一个双重循环遍历前两个数组。
	store := make(map[int]int)
	for _, value1 := range nums1 {
		for _, value2 := range nums2 {
			store[value1+value2]++
		}
	}
	count := 0
	// 第二个双重循环存储遍历后两个数组。
	for _, value3 := range nums3 {
		for _, value4 := range nums4 {
			if times, ok := store[-(value3 + value4)]; ok {
				count += times
			}
		}
	}

	return count
}

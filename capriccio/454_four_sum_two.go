package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// The key of map is to store the sum of two, the value is the number of times.
	store := make(map[int]int)
	for _, value1 := range nums1 {
		for _, value2 := range nums2 {
			store[value1+value2]++
		}
	}
	count := 0
	for _, value3 := range nums3 {
		for _, value4 := range nums4 {
			if times, ok := store[-(value3 + value4)]; ok {
				count += times
			}
		}
	}

	return count
}

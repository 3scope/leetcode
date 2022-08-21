package main

// 使用“map”进行去重。
func findPairs(nums []int, k int) int {
	// 存放每个“pair”中较小的那一个值，比如在“(1, 3)和(3, 5)“中，只存放1和3。
	result := make(map[int]bool)
	// 使用“map”去重。
	set := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if set[nums[i]-k] {
			result[nums[i]-k] = true
		}
		if set[nums[i]+k] {
			result[nums[i]] = true
		}
		set[nums[i]] = true
	}

	return len(result)
}

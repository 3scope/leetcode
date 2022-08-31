package main

import "sort"

// 首先可以将数组排序，之后通过排序后的下标，找出哪些数字比当前数字要小。
func smallerNumbersThanCurrent(nums []int) []int {
	// 首先对原数组排序，存放到一个新的数组，排序后数字在数组中的下标就表示有多少个数字比自己小。
	temp := make([]int, len(nums))
	copy(temp, nums)
	sort.Ints(temp)
	// 使用哈希表统计结果。
	count := make(map[int]int)
	// 从后往前统计结果，可以让相同的数字得到正确的结果，例如：[1, 2, 3, 3, 4, 5, 5, 6] => [0, 1, 2, 2, 4, 5, 5, 7]
	for i := len(temp) - 1; i >= 0; i-- {
		count[temp[i]] = i
	}
	// 按照原来的顺序输出结果。
	result := make([]int, len(nums))
	for i := 0; i < len(result); i++ {
		result[i] = count[nums[i]]
	}
	return result
}

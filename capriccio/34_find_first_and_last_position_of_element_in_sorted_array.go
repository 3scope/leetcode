package main

// 二分查找，首先找到“target”出现的最左边的位置，之后找到第一个比“target”大的数出现的位置。
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	leftIndex, ok := searchBinaryLeft(nums, target)
	if !ok {
		return []int{-1, -1}
	}
	// “target+1”如果不存在也会返回第一个大于“target”值所在的位置。
	rightIndex, _ := searchBinaryLeft(nums, target+1)

	return []int{leftIndex, rightIndex - 1}
}

// “bool”类型返回值代表值是否存在，“int”类型返回值代表位置的下标。
func searchBinaryLeft(nums []int, target int) (int, bool) {
	left := 0
	right := len(nums)
	isExisted := false
	// 只有当“left == right”才会跳出循环。
	for left < right {
		middle := left + (right-left)/2
		// 找到“target”出现的最小下标。
		if nums[middle] >= target {
			right = middle
		} else {
			left = middle + 1
		}
		if nums[middle] == target {
			isExisted = true
		}
	}
	return left, isExisted
}

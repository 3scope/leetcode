package main

import (
	"math/rand"
	"time"
)

func findKthLargest(nums []int, k int) int {
	result := 0
	quickSort(nums, 0, len(nums)-1, k, &result)
	return result
}

// 快速排序中如果排序的数字下标刚好是k，那么直接返回。
func quickSort(nums []int, left, right, k int, result *int) {
	// 终止条件。
	if left >= right {
		if left == k {
			*result = nums[left]
			return
		}
	}

	i, j := left, right
	// 枢纽选择为left。
	// pivot := nums[left]
	// 随机选择枢纽。
	rand.Seed(time.Now().UnixNano())
	offset := rand.Intn(right - left + 1)
	nums[left], nums[left+offset] = nums[left+offset], nums[left]
	pivot := nums[left]
	// 降序排列。
	for i < j {
		for i < j && nums[j] <= pivot {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= pivot {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = pivot

	if i == k-1 {
		*result = pivot
	} else if i > k-1 {
		// 如果所在位置大了，就去小一点的地方找。
		quickSort(nums, left, i-1, k, result)
	} else {
		quickSort(nums, i+1, right, k, result)
	}
}

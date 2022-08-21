package main

import "sort"

// 可以先将数组排序，这样就可以得到数对距离的最大和最小值，之后通过二分法求得第k小的数。
// 通过滑动窗口可以统计出有多少组数对小于当前的中间值“middle”。
func smallestDistancePair(nums []int, k int) int {
	sort.Ints(nums)
	// 二分查找的指针。
	left := 0
	right := nums[len(nums)-1] - nums[0]
	result := 0
	for left <= right {
		// 取数对最大与最小距离的中间值。
		middle := left + (right-left)/2
		// 统计有多少组数对小于当前的“middle”值。
		count := 0
		// 滑动窗口的指针。
		i, j := 0, 0
		for j < len(nums) {
			if nums[j]-nums[i] > middle {
				i++
			} else {
				count += j - i
				j++
			}
		}
		// 当“count”的值比k要小的时候，证明“middle”取值太小。
		if count < k {
			left = middle + 1
		} else if count >= k {
			// 当“count”的值比k要大或者等于k时，证明“middle”有可能就是目标值。
			result = middle
			right = middle - 1
		}
	}
	return result
}

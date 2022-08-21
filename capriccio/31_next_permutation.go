package main

import "sort"

// 使用字典序作为全排列顺序，那么第一个排列是顺序排列，最后一个是倒序排列。
func nextPermutation(nums []int) {
	if len(nums) == 1 {
		return
	}
	// 如果当前排列是“124653”，那么找它的下一个排列的方法是，在序列中从右至左找到第一个左邻小于右邻的数对。本例中第一个左邻小于右邻的数对是“46”，那么存储“4”的下标“2”为“i”，之后从右往左找第一个比“4”大的数，本例中是“5”，存储其下标为“j”。
	// 交换“i”与“j”，之后将排列中“i+1”到末尾的数设置为升序，最后的结果就是下一个排列。
	// 如果找不到，那么这是最后一个排列。
	// 从右到左寻找第一个正序对，即“i”的位置。
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	// 如果没找到，那么证明是最后一个排列。
	if i == -1 {
		sort.Ints(nums)
	}
	// 从右到左寻找第一个比“i”大的数，存在“j”中。
	j := len(nums) - 1
	for ; j > i; j-- {
		if nums[j] > nums[i] {
			break
		}
	}
	// 交换“i”和“j”，之后将“i”之后的数都设置为升序。
	nums[i], nums[j] = nums[j], nums[i]
	// 从“i+1”到末尾一定是降序。
	for left, right := i+1, len(nums)-1; left < right; left, right = left+1, right-1 {
		nums[left], nums[right] = nums[right], nums[left]
	}
}

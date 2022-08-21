package main

func removeElement(nums []int, val int) int {
	// 使用双指针法，可以将复杂度降到O(n)。
	slow, fast := 0, 0
	// “slow”指针存放最终的结果，“fast”指针用于判断数据是否需要删除。
	for fast < len(nums) {
		if nums[fast] == val {
			// 此时“fast”指向的值需要删除。
			fast++
		} else {
			// 此时“fast”指向的值需要保存。
			nums[slow] = nums[fast]
			slow++
			fast++
		}
	}

	nums = nums[:slow]
	return len(nums)
}

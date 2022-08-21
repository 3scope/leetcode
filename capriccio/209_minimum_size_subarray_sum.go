package main

func minSubArrayLen(target int, nums []int) int {
	// 使用滑动窗口，可以避免重复计算。
	// “sum”变量记录滑动窗口中所有元素的和，当“sum”的值大于“target”的时候，左指针需要右移，减少“sum”的值；当“sum”的值小于“target”的时候，右指针需要右移，增加“sum”的值。
	sum := 0
	// 不存在就返回0，因此初始化为0。
	result := 0
	// 左指针指向窗口的第一个值，右指针指向窗口的最后一个值的下一个。
	left := 0
	right := 0
	// 当右指针指向最后一个元素的下一个时，右指针就不能继续右移了。
	for right <= len(nums) {
		// 满足条件时，那么就尝试缩小窗口的值，直到不满足条件。
		if sum >= target {
			// 每次有新值需要记录下来，取其中最小的。
			if result == 0 || result > (right-left) {
				result = right - left
			}
			sum -= nums[left]
			left++
		} else if sum < target {
			if right < len(nums) {
				sum += nums[right]
			}
			right++
		}
	}

	return result
}

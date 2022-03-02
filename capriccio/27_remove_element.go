package main

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		// When the value of fast pointer equals to val, passed.
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
			fast++
		} else {
			fast++
		}
	}
	return slow
}

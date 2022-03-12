package main

func rotate(nums []int, k int) {
	length := len(nums)
	if length == 1 {
		return
	}
	k = k % length

	temp := make([]int, k)
	for i := 0; i < k; i++ {
		temp[i] = nums[length-k+i]
	}

	for slow, fast := length-k-1, length-1; slow >= 0; slow, fast = slow-1, fast-1 {
		nums[fast] = nums[slow]
	}

	for i := 0; i < k; i++ {
		nums[i] = temp[i]
	}
}

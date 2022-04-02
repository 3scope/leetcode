package main

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	result := make([]int, 2)
	result[0], result[1] = -1, -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			result[0] = mid
			right = mid - 1
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	if result[0] != -1 {
		for i := result[0]; i < len(nums) && nums[i] == target; i++ {
			result[1] = i
		}
	}

	return result
}

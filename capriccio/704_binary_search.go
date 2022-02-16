package capriccio

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left, right := 0, len(nums)-1
	middle := (left + right) / 2

	for left <= right {
		if nums[middle] == target {
			break
		} else if nums[middle] > target {
			right = middle - 1
			middle = (left + right) / 2
		} else {
			left = middle + 1
			middle = (left + right) / 2
		}
	}

	if left > right {
		return -1
	} else {
		return middle
	}
}

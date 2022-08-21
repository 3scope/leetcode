package main

// 使用双指针法，如果左右指针能够相遇，那么它就是一个山脉数组。
func validMountainArray(arr []int) bool {
	// 如果数组长度小于3，那么直接返回“false”。
	if len(arr) < 3 {
		return false
	}
	// 两个指针最后都会指向山峰。
	left := 0
	right := len(arr) - 1
	for left < right {
		// 如果两边都没法移动，那么跳出循环。
		if arr[left] >= arr[left+1] && arr[right] >= arr[right-1] {
			break
		}

		if arr[left] < arr[left+1] {
			left++
		}
		if arr[right] < arr[right-1] {
			right--
		}
	}
	// 如果数组是单调递增和单调递减的，那么返回“false”。
	if left == len(arr)-1 {
		return false
	} else if right == 0 {
		return false
	}
	// 如果最后两个指针没有重合，也就是没有同时指向山峰，那么返回“false”，否则返回“true”。
	return left == right
}

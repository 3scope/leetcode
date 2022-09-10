package main

// 时间复杂度要求为“O(n)”，最简单的实现方式是使用“map”，用空间换时间。
// 如果要求常数级的空间则使用原地“hash”。
// 可以采取这样的思路：把“1”这个数放到下标为“0”的位置，把“2”这个数放到下标为“1”的位置，然后再遍历一次数组，第一个“nums[i] - 1 != i”的数就是结果。
func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		if nums[i]-1 == i {
			i++
		} else if nums[i] < 1 || nums[i] > len(nums) || nums[i] == nums[nums[i]-1] {
			i++
		} else {
			swapInt(nums, i, nums[i]-1)
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i]-1 != i {
			return i + 1
		}
	}
	return len(nums) + 1
}

func swapInt(nums []int, i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

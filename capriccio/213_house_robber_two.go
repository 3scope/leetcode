package main

func robTwo(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	return max(robOne(nums, 0, len(nums)-2), robOne(nums, 1, len(nums)-1))
}

func robOne(nums []int, start, end int) int {
	if end == start {
		return nums[start]
	}
	max := func(x, y int) int {
		if x > y {
			return x
		} else {
			return y
		}
	}
	dp0 := nums[start]
	dp1 := max(nums[start], nums[start+1])
	result := max(dp0, dp1)
	for i := start + 2; i <= end; i++ {
		result = max(dp1, dp0+nums[i])
		dp0 = dp1
		dp1 = result
	}

	return result
}

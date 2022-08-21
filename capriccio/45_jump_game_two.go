package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cover := nums[0]
	nextCover := 0
	result := 0
	for i := 0; i <= cover; i++ {
		if i+nums[i] > nextCover {
			nextCover = i + nums[i]
		}
		if i == cover {
			cover = nextCover
			result++
			// 每次更新“cover”的值后，如果比数组长度要大，那么就可以在当前步数内走出去。
			if cover >= len(nums)-1 {
				break
			}
		}
	}

	return result
}

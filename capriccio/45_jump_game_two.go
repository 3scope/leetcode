package main

func jump(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	cover := 0
	nextCover := 0
	result := 0
	for i := 0; i <= cover; i++ {
		if i+nums[i] > nextCover {
			nextCover = i + nums[i]
		}
		if i == cover {
			cover = nextCover
			result++
			// Every time update the value of cover, need to check if the end point has been reached.
			if cover >= len(nums)-1 {
				break
			}
		}
	}

	return result
}

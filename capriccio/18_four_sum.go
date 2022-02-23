package capriccio

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	// Start iterating.
	result := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		for j := i + 1; j < len(nums)-2; j++ {
			if j > i+1 && nums[j-1] == nums[j] {
				continue
			}
			slow := j + 1
			fast := len(nums) - 1
			for slow < fast {
				if nums[i]+nums[j]+nums[slow]+nums[fast] > target {
					fast--
				} else if nums[i]+nums[j]+nums[slow]+nums[fast] < target {
					slow++
				} else {
					result = append(result, []int{nums[i], nums[j], nums[slow], nums[fast]})
					slow++
					fast--
					for slow < fast && nums[slow] == nums[slow-1] {
						slow++
					}
					for slow < fast && nums[fast] == nums[fast+1] {
						fast--
					}
				}
			}
		}
	}
	return result
}

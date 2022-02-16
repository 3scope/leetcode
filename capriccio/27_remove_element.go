package capriccio

func removeElement(nums []int, val int) int {
	slow := 0
	length := len(nums)
	for fast := 0; fast < length; fast++ {
		if val != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}

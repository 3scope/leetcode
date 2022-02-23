package capriccio

func intersection(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return nil
	}
	result := make([]int, 0)
	inter := make(map[int]int)
	// To remove duplicates.
	for _, value := range nums1 {
		inter[value] = 1
	}
	for _, value := range nums2 {
		if inter[value] == 1 {
			inter[value] = -1
			result = append(result, value)
		}
	}

	return result
}

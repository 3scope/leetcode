package main

func twoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for index, value := range nums {
		if i, ok := store[target-value]; ok {
			return []int{i, index}
		} else {
			store[value] = index
		}
	}
	return nil
}

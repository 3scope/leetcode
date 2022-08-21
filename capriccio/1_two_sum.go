package main

// 给出两数之和的下标，可以在第一次遍历到某个值的时候，将该值存储到“map”里，“key”为“num”，“value”为下标。
func twoSum(nums []int, target int) []int {
	store := make(map[int]int)

	for index, value := range nums {
		// 题目中给出每个样例只有一个正确答案，因此如果满足直接返回。
		if i, ok := store[target-value]; ok {
			return []int{i, index}
		} else {
			store[value] = index
		}
	}
	return nil
}

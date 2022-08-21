package main

func bestRotation(nums []int) int {
	// 差分数组，数组的下标代表“移动几步”。
	diff := make([]int, len(nums))
	// 根据每个元素需要移动的步数范围，确定差分数组。
	for i := 0; i < len(nums); i++ {
		// 先求一个元素理论上移动最多的步数，就是将该元素循环左移到“下标”和“值”相同的位置。
		// 即当前位置“i”减去目标位置“nums[i]”的值。
		// 而为了使得值为正，可以加上数组的长度，再取模。
		maxStep := (i - nums[i] + len(nums)) % len(nums)
		// 移动最少的步数就是移动到数组最后一位，因为最后一位下标值最大。
		minStep := (i - (len(nums) - 1) + len(nums)) % len(nums)
		// 差分数组的含义是，移动“i”步，那么该数可以满足条件，得到一分。
		diff[minStep]++
		if maxStep+1 < len(nums) {
			// 移动最大步数后，后面的步数不能得分。
			diff[maxStep+1]--
		}
		// 如果“minStep”比“maxStep”要大，那么意味着，移动的区间是不连续的，例如一个长度为8的数组，某个元素移动[0, 1]步可以，移动7步也可以，但是[2, 6]步不行，那么此“minStep”比“maxStep”要大。
		if minStep > maxStep {
			diff[0]++
		}
	}

	// 求前缀和，记录最大分数的最小移动步数。
	curScore := 0
	maxScore := 0
	index := -1
	for i := 0; i < len(diff); i++ {
		curScore += diff[i]
		if curScore > maxScore {
			maxScore = curScore
			index = i
		}
	}
	return index
}

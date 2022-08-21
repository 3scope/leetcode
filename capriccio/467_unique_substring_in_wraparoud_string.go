package main

func findSubstringInWraproundString(p string) int {
	// 字符串“s”是一个"abcdefghijklmnopqrstuvwxyz"的重复字符串，即相邻的两个元素的差值为1。
	// s[i] - 'a' = (s[i-1] + 1 - 'a') % 26
	// 那么同样的，如果“p”中的子串如果需要在“s”中出现，那么它的子串也会满足以上的条件。
	// 例如“p”为“abcdbcd”，那么对应的以“i”下标所在字符结尾的子串在“s”中的出现的次数为“count[i]”。
	// 子串出现的次数为[1, 1, 1, 1, 1, 1, 1]
	// 如果满足“s[i] = (s[i-1] + 1) % 26”，那么可以直接计算前缀和，如果不满足，那么就需求设置为1。
	// 因此 count = [1, 2, 3, 4, 1, 2 ,3]
	count := make([]int, len(p))
	// 得到“count”的值。
	count[0] = 1
	for i := 1; i < len(p); i++ {
		// 设置初始值。
		count[i] = 1
		if p[i]-'a' == (p[i-1]+1-'a')%26 {
			count[i] += count[i-1]
		} else {
			// 不满足条件，那么设置为1。
			count[i] = 1
		}
	}

	// 之后需要再次遍历一次“count”，因为“p”中会出现重复字符，只需要取其中的最大值。
	temp := make(map[byte]int)
	for i := 0; i < len(count); i++ {
		if temp[p[i]] < count[i] {
			temp[p[i]] = count[i]
		}
	}

	// 之后将所有结果加起来即可。
	result := 0
	for _, v := range temp {
		result += v
	}
	return result
}

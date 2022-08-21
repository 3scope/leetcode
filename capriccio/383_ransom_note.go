package main

// 通过“magazine”上的字符组成“ransomNote”，要求“ransomNote”中出现的字符，“magazine”中都出现过，并且数量要比“ransomNote”中的多。
func canConstruct(ransomNote string, magazine string) bool {
	// 存储“ransomNote”中各字符的数量。
	count := make(map[rune]int)
	for _, value := range ransomNote {
		count[value]++
	}
	// 遍历“magazine”，将相应的字符数量减一。
	for _, value := range magazine {
		count[value]--
	}
	// 如果最终“count”中出现大于1的“value”的话，那么证明“magazine”不满足要求。
	for _, value := range count {
		if value > 0 {
			return false
		}
	}
	return true
}

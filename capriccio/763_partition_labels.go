package main

func partitionLabels(s string) []int {
	position := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		if i > position[s[i]] {
			position[s[i]] = i
		}
	}

	result := make([]int, 0)
	begin := 0
	end := position[s[0]]
	for i := 0; i < len(s); i++ {
		if i == end {
			result = append(result, end-begin+1)
			begin = end + 1
			if i != len(s)-1 {
				end = position[s[i+1]]
			}
		} else if position[s[i]] > end {
			end = position[s[i]]
		}
	}

	return result
}

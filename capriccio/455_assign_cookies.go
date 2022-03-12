package main

import "sort"

func findContentChildren(g []int, s []int) int {
	if len(s) == 0 {
		return 0
	}
	sort.Ints(g)
	sort.Ints(s)
	result := 0
	for i, j := len(g)-1, len(s)-1; i >= 0 && j >= 0; i-- {
		if (s[j]) >= g[i] {
			result++
			j--
		}
	}

	return result
}

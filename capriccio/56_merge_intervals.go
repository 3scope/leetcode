package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	left := intervals[0][0]
	right := intervals[0][1]
	result := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= right {
			if intervals[i][0] < left {
				left = intervals[i][0]
			}
			right = intervals[i][1]
		} else {

		}
	}

	return result
}

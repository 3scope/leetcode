package main

import "sort"

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	right := intervals[0][1]
	result := 0
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= right {
			result++
			right = intervals[i][1]
		} else if intervals[i][0] < right {
			result++
		} else {
			right = intervals[i][1]
		}
	}

	return result
}

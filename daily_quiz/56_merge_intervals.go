package main

import "sort"

func merge(intervals [][]int) [][]int {
	result := make([][]int, 0)

	// First to sort.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	if intervals == nil || len(intervals) == 1 {
		return intervals
	}

	length := len(intervals) - 1
	for i := 0; i < length; i++ {
		if intervals[i][1] > intervals[i+1][1] {
			intervals[i+1] = intervals[i]
		} else if intervals[i][1] >= intervals[i+1][0] {
			intervals[i+1][0] = intervals[i][0]
		} else {
			result = append(result, intervals[i])
		}
	}
	result = append(result, intervals[length])

	return result
}

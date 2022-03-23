package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := make([][]int, 0)
	begin := intervals[0][0]
	end := intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][1] <= end {
			continue
		} else if intervals[i][0] <= end {
			end = intervals[i][1]
		} else {
			result = append(result, []int{begin, end})
			begin = intervals[i][0]
			end = intervals[i][1]
		}
	}
	result = append(result, []int{begin, end})

	return result
}

package main

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	right := points[0][1]
	result := 1
	// Sort by the left boundary, only three cases.
	for i := 1; i < len(points); i++ {
		if points[i][1] <= right {
			// Greedy, use the mininum length of balloons.
			right = points[i][1]
		} else if points[i][0] <= right {
			continue
		} else {
			right = points[i][1]
			result++
		}
	}

	return result
}

package main

import "math"

// 求三条边的边长，如果可以构成三角形，那么就是回旋镖。
func isBoomerang(points [][]int) bool {
	length1 := lengthOfTwoPoints(points[0], points[1])
	length2 := lengthOfTwoPoints(points[0], points[2])
	lenght3 := lengthOfTwoPoints(points[1], points[2])

	if length1+length2 > lenght3+1e-9 && length1+lenght3 > length2+1e-9 && length2+lenght3 > length1+1e-9 {
		return true
	} else {
		return false
	}
}

func lengthOfTwoPoints(point1, point2 []int) float64 {
	temp := (point1[0]-point2[0])*(point1[0]-point2[0]) + (point1[1]-point2[1])*(point1[1]-point2[1])
	return math.Sqrt(float64(temp))
}

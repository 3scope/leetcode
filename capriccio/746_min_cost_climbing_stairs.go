package main

func minCostClimbingStairs(cost []int) int {
	length := len(cost)
	min := func(x, y int) int {
		if x > y {
			return y
		} else {
			return x
		}
	}
	if length == 2 {
		return min(cost[0], cost[1])
	}
	dp0 := cost[0]
	dp1 := cost[1]
	dpi := 0
	for i := 2; i < length; i++ {
		dpi = min(dp0, dp1) + cost[i]
		dp0 = dp1
		dp1 = dpi
	}
	return min(dp0, dp1)
}

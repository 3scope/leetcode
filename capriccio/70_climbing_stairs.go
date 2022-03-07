package main

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}
	slow := 1
	fast := 2
	result := 0
	for i := 3; i <= n; i++ {
		result = slow + fast
		slow = fast
		fast = result
	}
	return result
}

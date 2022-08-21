package main

// 前缀和数组：原数组中的元素不变，求数组某个前缀区间的和。
// 例如：nums = [1,2,3,4,5] => preSum = [1,3,6,10,15] (或者是 [0,1,3,6,10,15])。
// func getPreSumArray(nums []int) {
// 	for i := 1; i < len(nums); i++ {
// 		nums[i] = nums[i] + nums[i-1]
// 	}
// }

// 差分数组：和前缀和数组互为逆运算，求两个元素之间的差值。
// 例如：arr = [1,2,3,4,5] => diff = [1,1,1,1,1]。
// func getDiffArray(nums []int) {
// 	for i := len(nums) - 1; i > 0; i-- {
// 		nums[i] = nums[i] - nums[i-1]
// 	}
// }

func corpFlightBookings(bookings [][]int, n int) []int {
	// 使用差分数组和前缀和数组，可以通过开始点和结束点统计整个区间的变化。
	result := make([]int, n+2)
	for i := 0; i < len(bookings); i++ {
		// 区间起始值。
		first := bookings[i][0]
		// “last”在前缀和中最后一个可以加上“seats”的地址，所以需要将“last+1”减去“seats”的值，那么之后的数据就不会受到影响。
		last := bookings[i][1]
		// 记录有多少个座位。
		seats := bookings[i][2]
		result[first] += seats
		result[last+1] -= seats
	}
	result = result[1 : len(result)-1]
	// 求前缀和。
	for i := 1; i < len(result); i++ {
		result[i] = result[i-1] + result[i]
	}

	return result
}

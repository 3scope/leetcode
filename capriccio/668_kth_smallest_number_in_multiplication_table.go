package main

func findKthNumber(m int, n int, k int) int {
	// 使用二分查找的方式，区间是 [left, right]。
	left := 1
	right := m * n
	for left <= right {
		// 这样写是为了防止整型溢出的情况。
		// 得到“middle”值，之后统计乘法表中有多少个数小于或者等于“middle”。
		middle := left + (right-left)/2
		count := countInMultiplication(middle, m, n)
		// 由于“count”的值是有多少个数小于或者等于“middle”，所以求第“k”小的数时，不是和“k-1”做比较。
		if count > k {
			right = middle - 1
		} else if count < k {
			left = middle + 1
		} else {
			// 相等时，查找左边区间的值，可能出现多个数字在乘法表中小于或者等于它的总个数为“k”，这时需要选择最小的。
			right = middle - 1
		}
	}
	// 最后left会指向最后一个
	return left
}

// 统计乘法表中有多少个数小于或者等于“n”。
func countInMultiplication(n int, rows, columns int) int {
	// 对每一行进行统计。
	count := 0
	for i := 1; i <= rows; i++ {
		// 每一行中小于或者等于“n”值的个数和行号有关。
		// “n / i”代表每一行中小于或者等于“n”值的个数，因为每一行的值都是行号“i”乘以列号“j”，那么“n / i”就可以知道多少列比“n”要小。
		eachRowCount := n / i
		// 如果列号更多的话，那么以每行最多比“n”小的个数为准。
		if columns > eachRowCount {
			count += eachRowCount
		} else {
			count += columns
		}
	}

	return count
}

package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	// 通过左上角的点和右下角的点不断逼近，不断的填充矩阵。
	leftTop := []int{0, 0}
	rightBottom := []int{n - 1, n - 1}

	// 从1开始填充。
	num := 1
	for leftTop[0] < rightBottom[0] && leftTop[1] < rightBottom[1] {
		// 上面的边，保证区间左闭右开。
		for i := leftTop[1]; i < rightBottom[1]; i++ {
			result[leftTop[0]][i] = num
			num++
		}
		// 右边的边。
		for i := leftTop[0]; i < rightBottom[0]; i++ {
			result[i][rightBottom[1]] = num
			num++
		}
		// 下面的边。
		for i := rightBottom[1]; i > leftTop[1]; i-- {
			result[rightBottom[0]][i] = num
			num++
		}
		// 左边的边。
		for i := rightBottom[0]; i > leftTop[0]; i-- {
			result[i][leftTop[1]] = num
			num++
		}
		leftTop[0]++
		leftTop[1]++
		rightBottom[0]--
		rightBottom[1]--
	}

	// 当“n”为奇数的时候，最中间的元素需要特殊处理。
	if n%2 == 1 {
		result[leftTop[0]][leftTop[1]] = num
	}

	return result
}

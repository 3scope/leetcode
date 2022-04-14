package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, n)
	}
	leftTop := []int{0, 0}
	rightBottom := []int{n - 1, n - 1}

	num := 1
	for leftTop[0] < rightBottom[0] && leftTop[1] < rightBottom[1] {
		// Top edge.
		for i := leftTop[1]; i < rightBottom[1]; i++ {
			result[leftTop[0]][i] = num
			num++
		}
		// Right edge.
		for i := leftTop[0]; i < rightBottom[0]; i++ {
			result[i][rightBottom[1]] = num
			num++
		}
		// Bottom edge.
		for i := rightBottom[1]; i > leftTop[1]; i-- {
			result[rightBottom[0]][i] = num
			num++
		}
		// Left edge.
		for i := rightBottom[0]; i > leftTop[0]; i-- {
			result[i][leftTop[1]] = num
		}
		leftTop[0]++
		leftTop[1]++
		rightBottom[0]--
		rightBottom[1]--
	}

	if n%2 == 1 {
		result[leftTop[0]][leftTop[1]] = num
	}

	return result
}

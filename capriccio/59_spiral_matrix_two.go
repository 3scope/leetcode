package main

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, n)
	}

	level := n / 2
	count := 1
	for i := 0; i < level; i++ {
		// Top edge.
		for j := i; j < n-i-1; j++ {
			result[i][j] = count
			count++
		}
		// Right edge.
		for j := i; j < n-i-1; j++ {
			result[j][n-i-1] = count
			count++
		}
		// Buttom edge.
		for j := n - i - 1; j > i; j-- {
			result[n-i-1][j] = count
			count++
		}
		// Left edge.
		for j := n - i - 1; j > i; j-- {
			result[j][i] = count
			count++
		}

	}

	if n%2 == 1 {
		result[n/2][n/2] = n * n
	}
	return result
}

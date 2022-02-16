package capriccio

func spiralOrder(matrix [][]int) []int {
	if matrix == nil {
		return nil
	}
	height := len(matrix)
	width := len(matrix[0])
	level := min(height, width) / 2
	result := make([]int, 0)

	for i := 0; i < level; i++ {
		// Top edge.
		for j := i; j < width-i-1; j++ {
			result = append(result, matrix[i][j])
		}
		// Right edge.
		for j := i; j < height-i-1; j++ {
			result = append(result, matrix[j][width-i-1])
		}
		// Buttom edge.
		for j := width - i - 1; j > i; j-- {
			result = append(result, matrix[height-i-1][j])
		}
		// Left edge.
		for j := height - i - 1; j > i; j-- {
			result = append(result, matrix[j][i])
		}
	}

	if min(height, width)%2 == 1 {
		i := level
		if width >= height {
			for j := i; j < width-i; j++ {
				result = append(result, matrix[i][j])
			}
		} else {
			for j := i; j < height-i; j++ {
				result = append(result, matrix[j][width-i-1])
			}
		}
	}

	return result
}

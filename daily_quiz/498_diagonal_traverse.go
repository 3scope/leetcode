package main

// 遍历元素的下标“[i, j]”之和“i+j”从0开始依次递增。
func findDiagonalOrder(mat [][]int) []int {
	rows := len(mat)
	columns := len(mat[0])
	// 存放结果。
	result := make([]int, 0, rows*columns)
	// 从下标之和为0的元素开始遍历，一直到下标之和为“rows+columns-2”为止。
	for sum := 0; sum <= rows+columns-2; sum++ {
		// 下标之和为偶数，从左下到右上角遍历。
		if sum%2 == 0 {
			for i := rows - 1; i >= 0; i-- {
				if i > sum {
					continue
				}
				j := sum - i
				if j > sum || j >= columns {
					continue
				}
				result = append(result, mat[i][j])
			}
		} else {
			// 下标之和为奇数，从右上到左下角遍历。
			for i := 0; i < rows; i++ {
				if i > sum {
					continue
				}
				j := sum - i
				if j > sum || j >= columns {
					continue
				}
				result = append(result, mat[i][j])
			}
		}
	}

	return result
}

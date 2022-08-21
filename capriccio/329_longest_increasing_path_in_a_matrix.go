package main

// 记忆化深度优先搜索。
func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 && len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	columns := len(matrix[0])
	memo := make([][]int, rows)
	for i := 0; i < rows; i++ {
		memo[i] = make([]int, columns)
	}
	result := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			subResult := dfsPath(matrix, memo, [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}, i, j)
			if result < subResult {
				result = subResult
			}
		}
	}
	return result
}

// 记忆化搜索矩阵“memo”代表如果以当前节点“[i, j]”为终点，最长路径为“memo[i, j]”。
func dfsPath(matrix, memo, directions [][]int, i, j int) int {

	if memo[i][j] != 0 {
		return memo[i][j]
	}
	// 初始化为1。
	memo[i][j] = 1
	for _, value := range directions {
		newRow := value[0] + i
		newColumn := value[1] + j
		if newRow >= 0 && newRow < len(matrix) && newColumn >= 0 && newColumn < len(matrix[0]) && matrix[newRow][newColumn] > matrix[i][j] {
			// 取路径最大值。
			neb := dfsPath(matrix, memo, directions, newRow, newColumn)
			if neb+1 > memo[i][j] {
				memo[i][j] = neb + 1
			}
		}
	}
	return memo[i][j]
}

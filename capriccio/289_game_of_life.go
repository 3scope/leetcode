package main

func gameOfLife(board [][]int) {
	// 由于状态是由上一个状态推导的，因此需要重新开辟空间存储状态。
	temp := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		temp[i] = make([]int, len(board[i]))
	}
	// 方向数组，用于得到周围的值。
	directionRow := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	directionColumn := []int{0, 0, -1, 1, -1, 1, 1, -1}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			nums := 0
			// 遍历周围八个格子。
			for k := 0; k < 8; k++ {
				row := i + directionRow[k]
				column := j + directionColumn[k]
				if row >= 0 && row < len(board) && column >= 0 && column < len(board[i]) {
					nums += board[row][column]
				}
			}
			// 根据不同的状态推导下一个状态。
			switch {
			case nums < 2:
				temp[i][j] = 0
			case nums > 3:
				temp[i][j] = 0
			case nums == 3 && board[i][j] == 0:
				temp[i][j] = 1
			default:
				temp[i][j] = board[i][j]
			}
		}
	}
	copy(board, temp)
}

package main

func gameOfLife(board [][]int) {
	// Save the latest status.
	temp := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		temp[i] = make([]int, len(board[i]))
	}
	// Direction array.
	directionX := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	directionY := []int{0, 0, -1, 1, -1, 1, 1, -1}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			nums := 0
			for k := 0; k < 8; k++ {
				x := i + directionX[k]
				y := j + directionY[k]
				if x >= 0 && x < len(board) && y >= 0 && y < len(board[i]) {
					nums += board[x][y]
				}
			}
			// The variable "nums" store the number of viable cells.
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

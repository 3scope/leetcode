package main

func solveSudoku(board [][]byte) {
	solveSudokuBacktracking(board)
}

func solveSudokuBacktracking(board [][]byte) bool {
	height := len(board)
	width := len((board)[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] != '.' {
				continue
			}
			var k byte
			for k = '1'; k <= '9'; k++ {
				if isNumberValid(i, j, k, board) {
					board[i][j] = k
					if solveSudokuBacktracking(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}

	// When loop is over, return true.
	return true
}

func isNumberValid(row, column int, k byte, board [][]byte) bool {
	// For each row.
	for i := 0; i < 9; i++ {
		if board[row][i] == k {
			return false
		}
	}
	// For each column.
	for i := 0; i < 9; i++ {
		if board[i][column] == k {
			return false
		}
	}
	// For each 3 * 3 square.
	startRow := (row / 3) * 3
	startColumn := (column / 3) * 3
	for i := startRow; i < startRow+3; i++ {
		for j := startColumn; j < startColumn+3; j++ {
			if board[i][j] == k {
				return false
			}
		}
	}

	return true
}

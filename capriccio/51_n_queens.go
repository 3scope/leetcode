package main

func solveNQueens(n int) [][]string {
	result := make([][]string, 0)
	board := make([][]byte, n)
	for i := 0; i < n; i++ {
		board[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			board[i][j] = '.'
		}
	}
	solveNQueensBacktracking(n, &board, &result, 0)

	return result
}

func solveNQueensBacktracking(n int, board *[][]byte, result *[][]string, row int) {
	if n == row {
		temp := make([]string, n)
		for i := 0; i < n; i++ {
			temp[i] = string((*board)[i])
		}
		*result = append(*result, temp)
		return
	}

	if row > n {
		return
	}

	for i := 0; i < n; i++ {
		if isQueenValid(board, row, i) {
			(*board)[row][i] = 'Q'
			solveNQueensBacktracking(n, board, result, row+1)
			(*board)[row][i] = '.'
		}
	}
}

func isQueenValid(board *[][]byte, row, column int) bool {
	// Cannot be same column.
	for i := 0; i < row; i++ {
		if (*board)[i][column] == 'Q' {
			return false
		}
	}
	// Cannot be diagonal.
	for i, j := row-1, column-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if (*board)[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, column+1; i >= 0 && j < len(*board); i, j = i-1, j+1 {
		if (*board)[i][j] == 'Q' {
			return false
		}
	}

	return true
}

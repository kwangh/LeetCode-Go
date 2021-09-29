package leetcode

func solveSudoku(board [][]byte) {
	backtrack(board, 0, 0)
}

func backtrack(board [][]byte, i, j int) bool {
	if i == 9 {
		return true
	}
	if j == 9 {
		return backtrack(board, i+1, 0)
	}
	if board[i][j] != '.' {
		return backtrack(board, i, j+1)
	}
	for c := byte('1'); c <= byte('9'); c++ {
		if !isValid(board, i, j, c) {
			continue
		}
		board[i][j] = c
		if backtrack(board, i, j+1) {
			return true
		}
		board[i][j] = '.'
	}
	return false
}

func isValid(board [][]byte, row, col int, c byte) bool {
	cellRow, cellCol := 3*(row/3), 3*(col/3)
	for i := 0; i < 9; i++ {
		if board[i][col] == c || board[row][i] == c || board[cellRow+i/3][cellCol+i%3] == c {
			return false
		}
	}
	return true
}

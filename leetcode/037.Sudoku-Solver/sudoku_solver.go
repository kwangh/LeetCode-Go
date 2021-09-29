package leetcode

func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] != '.' {
				continue
			}
			for c := byte('1'); c <= byte('9'); c++ {
				if isValid(board, i, j, c) {
					board[i][j] = c

					if solve(board) {
						return true
					}
					board[i][j] = '.'
				}
			}
			return false
		}
	}
	return true
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

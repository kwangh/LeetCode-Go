package leetcode

func solve(board [][]byte) {
	for i := 0; i < len(board); i++ {
		check(board, i, 0)
		if len(board[0]) > 1 {
			check(board, i, len(board[0])-1)
		}
	}

	for j := 1; j < len(board[0])-1; j++ {
		check(board, 0, j)
		if len(board) > 1 {
			check(board, len(board)-1, j)
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
			if board[i][j] == '1' {
				board[i][j] = 'O'
			}
		}
	}
}

func check(board [][]byte, row, col int) {
	if board[row][col] == 'O' {
		board[row][col] = '1'
		if col+1 < len(board[0]) {
			check(board, row, col+1)
		}
		if col > 1 {
			check(board, row, col-1)
		}
		if row+1 < len(board) {
			check(board, row+1, col)
		}
		if row > 1 {
			check(board, row-1, col)
		}
	}
}

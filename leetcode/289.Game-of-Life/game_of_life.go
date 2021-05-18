package leetcode

func gameOfLife(board [][]int) {
	b := make([][]int, len(board))
	for i := range b {
		b[i] = make([]int, len(board[0]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			cnt := 0
			if i-1 >= 0 {
				if j-1 >= 0 && board[i-1][j-1] == 1 {
					cnt++
				}
				if board[i-1][j] == 1 {
					cnt++
				}

				if j+1 < len(board[0]) {
					if board[i-1][j+1] == 1 {
						cnt++
					}
				}
			}
			if j-1 >= 0 && board[i][j-1] == 1 {
				cnt++
			}
			if j+1 < len(board[0]) && board[i][j+1] == 1 {
				cnt++
			}
			if i+1 < len(board) {
				if j-1 >= 0 && board[i+1][j-1] == 1 {
					cnt++
				}
				if board[i+1][j] == 1 {
					cnt++
				}

				if j+1 < len(board[0]) && board[i+1][j+1] == 1 {
					cnt++
				}
			}

			if cnt == 2 && board[i][j] == 1 {
				b[i][j] = 1
			}
			if cnt == 3 {
				b[i][j] = 1
			}
		}
	}
	copy(board, b)
}

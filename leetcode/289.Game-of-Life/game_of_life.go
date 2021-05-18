package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Lives(board [][]int, i, j int) int {
	var lives int
	for x := max(i-1, 0); x <= min(i+1, len(board)-1); x++ {
		for y := max(j-1, 0); y <= min(j+1, len(board[0])-1); y++ {
			lives += board[x][y] & 1
		}
	}
	lives -= board[i][j] & 1
	return lives
}

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			lives := Lives(board, i, j)

			if (lives == 2 || lives == 3) && board[i][j] == 1 {
				board[i][j] = 3
			}
			if lives == 3 && board[i][j] == 0 {
				board[i][j] = 2
			}
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

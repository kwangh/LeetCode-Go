package leetcode

func isValidSudoku(board [][]byte) bool {
	h, v, box := make([][]bool, 9), make([][]bool, 9), make([][]bool, 9)
	for i := 0; i < len(h); i++ {
		h[i] = make([]bool, 9)
		v[i] = make([]bool, 9)
		box[i] = make([]bool, 9)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '.' {
				continue
			}
			cur := board[i][j] - '1'
			if h[i][cur] {
				return false
			} else {
				h[i][cur] = true
			}

			if v[j][cur] {
				return false
			} else {
				v[j][cur] = true
			}

			boxIndex := j/3 + (i/3)*3
			if box[boxIndex][cur] {
				return false
			} else {
				box[boxIndex][cur] = true
			}
		}
	}
	return true
}

package leetcode

func solveNQueens(n int) [][]string {
	puzzle := make([][]byte, n)
	for i := range puzzle {
		puzzle[i] = make([]byte, n)
		for j := 0; j < n; j++ {
			puzzle[i][j] = '.'
		}
	}
	var res [][]string
	backtrack(n, 0, puzzle, &res)
	return res
}

func backtrack(n, row int, puzzle [][]byte, res *[][]string) {
	if n == row {
		var tmp []string
		for i := range puzzle {
			tmp = append(tmp, string(puzzle[i]))
		}
		*res = append(*res, tmp)
		return
	}

	for col := 0; col < n; col++ {
		if isValid(row, col, n, puzzle) {
			puzzle[row][col] = 'Q'
			backtrack(n, row+1, puzzle, res)
			puzzle[row][col] = '.'
		}
	}
}

func isValid(row, col, n int, puzzle [][]byte) bool {
	for i := 0; i < row; i++ {
		if puzzle[i][col] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if puzzle[i][j] == 'Q' {
			return false
		}
	}
	for i, j := row-1, col+1; i >= 0 && j < n; i, j = i-1, j+1 {
		if puzzle[i][j] == 'Q' {
			return false
		}
	}
	return true
}

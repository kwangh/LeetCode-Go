package leetcode

func totalNQueens(n int) int {
	puzzle := make([][]int, n)
	for i := range puzzle {
		puzzle[i] = make([]int, n)
	}
	var res int
	backtrack(puzzle, 0, &res)
	return res
}

func backtrack(puzzle [][]int, row int, cnt *int) {
	if row == len(puzzle) {
		*cnt++
		return
	}

	for col := 0; col < len(puzzle); col++ {
		if isValid(puzzle, row, col) {
			puzzle[row][col] = 1
			backtrack(puzzle, row+1, cnt)
			puzzle[row][col] = 0
		}
	}
}

func isValid(puzzle [][]int, row, col int) bool {
	for i := 0; i < row; i++ {
		if puzzle[i][col] == 1 {
			return false
		}
	}
	for i := 1; i <= row; i++ {
		if col-i >= 0 && puzzle[row-i][col-i] == 1 {
			return false
		}
		if col+i < len(puzzle) && puzzle[row-i][col+i] == 1 {
			return false
		}
	}
	return true
}

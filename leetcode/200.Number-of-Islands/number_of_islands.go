package leetcode

func backtrack(grid [][]byte, row, col int) {
	grid[row][col] = '2'
	if row > 0 && grid[row-1][col] == '1' {
		backtrack(grid, row-1, col)
	}
	if row < len(grid)-1 && grid[row+1][col] == '1' {
		backtrack(grid, row+1, col)
	}
	if col > 0 && grid[row][col-1] == '1' {
		backtrack(grid, row, col-1)
	}
	if col < len(grid[0])-1 && grid[row][col+1] == '1' {
		backtrack(grid, row, col+1)
	}
}

func numIslands(grid [][]byte) int {
	var res int
	for row := range grid {
		for col, v := range grid[row] {
			if v == '1' {
				backtrack(grid, row, col)
				res++
			}
		}
	}

	return res
}

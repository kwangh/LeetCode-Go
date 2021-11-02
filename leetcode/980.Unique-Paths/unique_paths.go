package leetcode

func uniquePathsIII(grid [][]int) int {
	var res, zeros, x, y int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				zeros++
			} else if grid[i][j] == 1 {
				x, y = i, j
			}
		}
	}
	dfs(grid, x, y, zeros+1, &res)
	return res
}

func dfs(grid [][]int, row, col, zeros int, res *int) {
	if row == -1 || col == -1 || row == len(grid) || col == len(grid[0]) || grid[row][col] == -1 {
		return
	}

	if grid[row][col] == 2 {
		if zeros == 0 {
			*res++
		}
		return
	}

	grid[row][col] = -1
	dfs(grid, row-1, col, zeros-1, res)
	dfs(grid, row, col-1, zeros-1, res)
	dfs(grid, row+1, col, zeros-1, res)
	dfs(grid, row, col+1, zeros-1, res)
	grid[row][col] = 0
}

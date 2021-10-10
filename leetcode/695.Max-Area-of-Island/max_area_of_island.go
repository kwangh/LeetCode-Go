package leetcode

func maxAreaOfIsland(grid [][]int) int {
	var max int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				var cur int
				dfs(grid, i, j, &max, &cur)
			}
		}
	}
	return max
}

func dfs(grid [][]int, row, col int, max, cur *int) {
	if row < 0 || col < 0 || row == len(grid) || col == len(grid[0]) || grid[row][col] == 0 {
		return
	}

	grid[row][col] = 0
	*cur++
	if *cur > *max {
		*max = *cur
	}

	dfs(grid, row-1, col, max, cur)
	dfs(grid, row, col-1, max, cur)
	dfs(grid, row+1, col, max, cur)
	dfs(grid, row, col+1, max, cur)
}

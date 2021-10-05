package leetcode

func islandPerimeter(grid [][]int) int {
	var res, repeat int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				res++
				if i != 0 && grid[i-1][j] == 1 {
					repeat++
				}
				if j != 0 && grid[i][j-1] == 1 {
					repeat++
				}
			}
		}
	}
	return 4*res - 2*repeat
}

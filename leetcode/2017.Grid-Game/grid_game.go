package leetcode

func Min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func gridGame(grid [][]int) int64 {
	var top, bottom int64
	for i := 1; i < len(grid[0]); i++ {
		top += int64(grid[0][i])
	}
	min := top

	for i := 1; i < len(grid[0]); i++ {
		top -= int64(grid[0][i])
		bottom += int64(grid[1][i-1])
		min = Min(Max(top, bottom), min)
	}
	return min
}

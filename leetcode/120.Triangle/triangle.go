package leetcode

func minimumTotal(triangle [][]int) int {
	rows := len(triangle)
	sums := triangle[rows-1]
	for row := rows - 2; row >= 0; row-- {
		for col := range triangle[row] {
			sums[col] = min(sums[col], sums[col+1]) + triangle[row][col]
		}
	}
	return sums[0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

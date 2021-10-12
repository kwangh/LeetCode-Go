package leetcode

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	var fresh int
	var q [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				fresh++
			} else if grid[i][j] == 2 {
				q = append(q, []int{i, j})
			}
		}
	}
	if fresh == 0 {
		return 0
	}
	dir := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var cnt int
	for len(q) != 0 {
		size := len(q)
		cnt++
		for i := 0; i < size; i++ {
			cur := q[i]
			for _, d := range dir {
				row, col := cur[0]+d[0], cur[1]+d[1]
				if row == -1 || col == -1 || row == m || col == n || grid[row][col] == 0 || grid[row][col] == 2 {
					continue
				}
				fresh--
				grid[row][col] = 2
				q = append(q, []int{row, col})
			}
		}
		q = q[size:]
	}

	if fresh == 0 {
		return cnt - 1
	}
	return -1
}

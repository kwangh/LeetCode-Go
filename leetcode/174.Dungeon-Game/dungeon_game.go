package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calculateMinimumHP(dungeon [][]int) int {
	m, n := len(dungeon), len(dungeon[0])
	hp := make([][]int, m+1)
	for i := range hp {
		hp[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			hp[i][j] = 200000
		}
	}
	hp[m][n-1], hp[m-1][n] = 1, 1
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			need := min(hp[i+1][j], hp[i][j+1]) - dungeon[i][j]
			if need <= 0 {
				hp[i][j] = 1
			} else {
				hp[i][j] = need
			}
		}
	}
	return hp[0][0]
}

package leetcode

func updateMatrix(mat [][]int) [][]int {
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				mat[i][j] = 10000
			}
		}
	}

	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				continue
			}
			if i > 0 {
				mat[i][j] = Min(mat[i][j], mat[i-1][j]+1)
			}
			if j > 0 {
				mat[i][j] = Min(mat[i][j], mat[i][j-1]+1)
			}
		}
	}
	for i := len(mat) - 1; i >= 0; i-- {
		for j := len(mat[0]) - 1; j >= 0; j-- {
			if mat[i][j] == 0 {
				continue
			}
			if i < len(mat)-1 {
				mat[i][j] = Min(mat[i][j], mat[i+1][j]+1)
			}
			if j < len(mat[0])-1 {
				mat[i][j] = Min(mat[i][j], mat[i][j+1]+1)
			}
		}
	}
	return mat
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func updateMatrixDFS(mat [][]int) [][]int {
	q := [][]int{}
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 0 {
				q = append(q, []int{i, j})
			} else {
				mat[i][j] = 10000
			}
		}
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for len(q) != 0 {
		cell := q[0]
		q = q[1:]
		for _, d := range dirs {
			newRow, newCol := cell[0]+d[0], cell[1]+d[1]
			if newRow == -1 || newCol == -1 || newRow == len(mat) || newCol == len(mat[0]) || mat[newRow][newCol] <= mat[cell[0]][cell[1]]+1 {
				continue
			}
			mat[newRow][newCol] = mat[cell[0]][cell[1]] + 1
			q = append(q, []int{newRow, newCol})
		}
	}
	return mat
}

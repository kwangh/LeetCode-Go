package leetcode

func search(matrix [][]int, v [][]bool, target, row, col int) bool {
	if row == len(matrix) || col == len(matrix[0]) {
		return false
	}
	if v[row][col] {
		return false
	}
	v[row][col] = true

	if matrix[row][col] == target {
		return true
	}

	if matrix[row][col] < target {
		return search(matrix, v, target, row+1, col) || search(matrix, v, target, row, col+1)
	}

	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	v := make([][]bool, len(matrix))
	for i := range v {
		v[i] = make([]bool, len(matrix[0]))
	}
	return search(matrix, v, target, 0, 0)
}

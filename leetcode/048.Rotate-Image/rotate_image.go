package leetcode

func rotate(matrix [][]int) {
	for row := 0; row < len(matrix)/2; row++ {
		for col := row; col < len(matrix)-1-row; col++ {
			matrix[row][col], matrix[col][len(matrix)-1-row] = matrix[col][len(matrix)-1-row], matrix[row][col]
			matrix[row][col], matrix[len(matrix)-1-row][len(matrix)-1-col] = matrix[len(matrix)-1-row][len(matrix)-1-col], matrix[row][col]
			matrix[row][col], matrix[len(matrix)-1-col][row] = matrix[len(matrix)-1-col][row], matrix[row][col]
		}
	}
}

func rotateV1(matrix [][]int) {
	for i := 0; i < len(matrix)-1; i++ {
		for j := i + 1; j < len(matrix[i]); j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i])/2; j++ {
			matrix[i][j], matrix[i][len(matrix[j])-j-1] = matrix[i][len(matrix[j])-j-1], matrix[i][j]
		}
	}
}

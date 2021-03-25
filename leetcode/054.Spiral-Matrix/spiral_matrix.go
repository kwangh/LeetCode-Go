package leetcode

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	startRow, endRow, startCol, endCol := 0, len(matrix)-1, 0, len(matrix[0])-1
	dir := 0
	for startRow <= endRow && startCol <= endCol {
		switch dir {
		case 0: //Right
			for col := startCol; col <= endCol; col++ {
				res = append(res, matrix[startRow][col])
			}
			startRow++
		case 1: //Down
			for row := startRow; row <= endRow; row++ {
				res = append(res, matrix[row][endCol])
			}
			endCol--
		case 2: //Left
			for col := endCol; col >= startCol; col-- {
				res = append(res, matrix[endRow][col])
			}
			endRow--
		case 3: //Up
			for row := endRow; row >= startRow; row-- {
				res = append(res, matrix[row][startCol])
			}
			startCol++
		}
		dir = (dir + 1) % 4
	}
	return res
}

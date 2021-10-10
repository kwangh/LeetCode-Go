package leetcode

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	if image[sr][sc] != newColor {
		dfs(image, sr, sc, newColor, image[sr][sc])
	}
	return image
}

func dfs(image [][]int, row, col, newColor, startNum int) {
	if row == len(image) ||
		col == len(image[0]) ||
		row < 0 ||
		col < 0 ||
		image[row][col] == newColor ||
		image[row][col] != startNum {
		return
	}

	image[row][col] = newColor
	dfs(image, row-1, col, newColor, startNum)
	dfs(image, row, col-1, newColor, startNum)
	dfs(image, row+1, col, newColor, startNum)
	dfs(image, row, col+1, newColor, startNum)
}

package leetcode

func existNextChar(visited [][]bool, board [][]byte, i, j int, word string, c int) bool {
	if c == len(word)-1 {
		return board[i][j] == word[c]
	}

	if word[c] == board[i][j] {
		visited[i][j] = true
		if i+1 < len(board) && !visited[i+1][j] && existNextChar(visited, board, i+1, j, word, c+1) {
			return true
		}
		if i-1 >= 0 && !visited[i-1][j] && existNextChar(visited, board, i-1, j, word, c+1) {
			return true
		}
		if j+1 < len(board[0]) && !visited[i][j+1] && existNextChar(visited, board, i, j+1, word, c+1) {
			return true
		}
		if j-1 >= 0 && !visited[i][j-1] && existNextChar(visited, board, i, j-1, word, c+1) {
			return true
		}
		visited[i][j] = false
	}

	return false
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := range board {
		for j := range board[0] {
			if existNextChar(visited, board, i, j, word, 0) {
				return true
			}
		}
	}
	return false
}

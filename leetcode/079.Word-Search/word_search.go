package leetcode

func start(visited *[][]bool, board [][]byte, i, j int, word string) bool {
	if word == "" {
		return true
	}

	if i >= 0 && i < len(board) && j >= 0 && j < len(board[0]) {
		if !(*visited)[i][j] && word[0] == board[i][j] {
			(*visited)[i][j] = true
			if start(visited, board, i+1, j, word[1:]) ||
				start(visited, board, i-1, j, word[1:]) ||
				start(visited, board, i, j+1, word[1:]) ||
				start(visited, board, i, j-1, word[1:]) {
				return true
			} else {
				(*visited)[i][j] = false
			}
		}
	}

	return false
}

func exist(board [][]byte, word string) bool {
	visited := make([][]bool, len(board))
	for i := range visited {
		visited[i] = make([]bool, len(board[0]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] {
				if start(&visited, board, i, j, word) {
					return true
				}
				for k := 0; k < len(board); k++ {
					for l := 0; l < len(board[0]); l++ {
						visited[k][l] = false
					}
				}
			}
		}
	}
	return false
}

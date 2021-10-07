package leetcode

func dfs(board [][]byte, i, j int, word string) bool {
	if len(word) == 1 {
		return true
	}

	board[i][j] = '0'
	if i-1 >= 0 && board[i-1][j] == word[1] && dfs(board, i-1, j, word[1:]) {
		return true
	}
	if j-1 >= 0 && board[i][j-1] == word[1] && dfs(board, i, j-1, word[1:]) {
		return true
	}
	if i+1 < len(board) && board[i+1][j] == word[1] && dfs(board, i+1, j, word[1:]) {
		return true
	}
	if j+1 < len(board[0]) && board[i][j+1] == word[1] && dfs(board, i, j+1, word[1:]) {
		return true
	}
	board[i][j] = word[0]
	return false
}

func exist(board [][]byte, word string) bool {
	for i := range board {
		for j := range board[0] {
			if board[i][j] == word[0] && dfs(board, i, j, word) {
				return true
			}
		}
	}
	return false
}

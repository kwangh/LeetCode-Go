package leetcode

func findWords(board [][]byte, words []string) []string {
	root := construct(words)
	res := []string{}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, i, j, root, &res)
		}
	}
	return res
}

func dfs(board [][]byte, row, col int, t *Trie, res *[]string) {
	c := board[row][col]
	if c == '0' || t.Next[c-'a'] == nil {
		return
	}
	t = t.Next[c-'a']
	if t.Word != "" {
		*res = append(*res, t.Word)
		t.Word = ""
	}

	board[row][col] = '0'
	if row > 0 {
		dfs(board, row-1, col, t, res)
	}
	if col > 0 {
		dfs(board, row, col-1, t, res)
	}
	if row < len(board)-1 {
		dfs(board, row+1, col, t, res)
	}
	if col < len(board[0])-1 {
		dfs(board, row, col+1, t, res)
	}
	board[row][col] = c
}

type Trie struct {
	Next [26]*Trie
	Word string
}

func construct(words []string) *Trie {
	root := &Trie{}
	for _, word := range words {
		p := root
		for _, c := range word {
			if p.Next[c-'a'] == nil {
				p.Next[c-'a'] = &Trie{}
			}
			p = p.Next[c-'a']
		}
		p.Word = word
	}
	return root
}

package leetcode

func palindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func dfs(res *[][]string, cur []string, s string, start int) {
	if start == len(s) {
		*res = append(*res, append([]string{}, cur...))
	}
	for end := start; end < len(s); end++ {
		if palindrome(s[start : end+1]) {
			cur = append(cur, s[start:end+1])
			dfs(res, cur, s, end+1)
			cur = cur[:len(cur)-1]
		}
	}
}

func partition(s string) [][]string {
	var res [][]string
	dfs(&res, []string{}, s, 0)
	return res
}

package leetcode

func dfs(res *[][]string, cur []string, s string, start int, dp [][]bool) {
	if start == len(s) {
		*res = append(*res, append([]string{}, cur...))
		return
	}
	for end := start; end < len(s); end++ {
		if s[start] == s[end] && (end-start <= 2 || dp[start+1][end-1]) {
			dp[start][end] = true
			cur = append(cur, s[start:end+1])
			dfs(res, cur, s, end+1, dp)
			cur = cur[:len(cur)-1]
		}
	}
}

func partition(s string) [][]string {
	var res [][]string
	dp := make([][]bool, len(s))
	for i := range dp {
		dp[i] = make([]bool, len(s))
	}
	dfs(&res, []string{}, s, 0, dp)
	return res
}

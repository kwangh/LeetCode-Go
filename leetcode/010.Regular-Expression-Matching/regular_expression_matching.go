package leetcode

// IsMatchRecur https://leetcode.com/problems/regular-expression-matching/
func IsMatchRecur(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}
	sb, pb := []byte(s), []byte(p)
	firstMatch := (len(s) != 0 && (pb[0] == sb[0] || pb[0] == '.'))

	if len(p) >= 2 && pb[1] == '*' {
		return (IsMatchRecur(s, string(pb[2:])) ||
			(firstMatch && IsMatchRecur(string(sb[1:]), p)))
	}
	return firstMatch && IsMatchRecur(string(sb[1:]), string(pb[1:]))
}

// IsMatch regular expression matching
func IsMatch(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true
	sb, pb := []byte(s), []byte(p)
	for i := len(s); i >= 0; i-- {
		for j := len(p) - 1; j >= 0; j-- {
			firstMatch := (i < len(s) && (pb[j] == sb[i] || pb[j] == '.'))
			if j+1 < len(p) && pb[j+1] == '*' {
				dp[i][j] = dp[i][j+2] || (firstMatch && dp[i+1][j])
			} else {
				dp[i][j] = firstMatch && dp[i+1][j+1]
			}
		}
	}

	return dp[0][0]
}

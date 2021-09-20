package leetcode

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := range dp {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 0; i < len(s); i++ {
		dp[0][i] = 1
	}

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if s[j] == t[i] {
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i+1][j]
			}
		}
	}
	return dp[len(t)][len(s)]
}

package leetcode

func wordBreak(s string, wordDict []string) bool {
	m := make(map[string]struct{}, len(wordDict))
	for _, v := range wordDict {
		m[v] = struct{}{}
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				if _, ok := m[s[j:i]]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len(s)]
}

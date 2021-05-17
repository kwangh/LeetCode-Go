package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

var dp = [10001]int{}

func numSquares(n int) int {
	var dfs func(i int) int
	dfs = func(i int) int {
		if i == 0 {
			return 0
		}

		if dp[i] > 0 {
			return dp[i]
		}

		res := i
		for j := 1; j*j <= i; j++ {
			res = min(res, dfs(i-j*j))
		}

		res++
		dp[i] = res

		return res
	}

	return dfs(n)
}

package leetcode

func countBits(n int) []int {
	m := make([]int, n+1)
	for i := 1; i <= n; i++ {
		m[i] = m[i>>1] + (i & 1)
	}
	return m
}

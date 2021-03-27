package leetcode

func uniquePaths(m int, n int) int {
	paths := make([]int, m*n)
	for i := 0; i < n; i++ {
		paths[i] = 1
	}
	for i := 0; i < m; i++ {
		paths[i*n] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			paths[i*n+j] = paths[(i-1)*n+j] + paths[i*n+j-1]
		}
	}

	return paths[m*n-1]
}

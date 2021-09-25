package leetcode

func generateMatrix(n int) [][]int {
	top, down, left, right := 0, n-1, 0, n-1
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}
	num := 1
	for top <= down && left <= right {
		for i := left; i <= right; i++ {
			res[top][i] = num
			num++
		}
		top++

		for i := top; i <= down; i++ {
			res[i][right] = num
			num++
		}
		right--

		for i := right; i >= left; i-- {
			res[down][i] = num
			num++
		}
		down--

		for i := down; i >= top; i-- {
			res[i][left] = num
			num++
		}
		left++
	}
	return res
}

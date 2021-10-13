package leetcode

func combine(n int, k int) [][]int {
	var res [][]int
	mutate(&res, make([]int, 0, k), 1, n, k)
	return res
}

func mutate(res *[][]int, cur []int, start, n, k int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*res = append(*res, tmp)
		return
	}
	for i := start; i <= n-k+1; i++ {
		cur = append(cur, i)
		mutate(res, cur, i+1, n, k-1)
		cur = cur[:len(cur)-1]
	}
}

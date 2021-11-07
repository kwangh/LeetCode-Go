package leetcode

func minimizedMaximum(n int, quantities []int) int {
	l, r := 1, 100000
	for l < r {
		m, sum := (l+r)/2, 0
		for _, q := range quantities {
			sum += (q + m - 1) / m
		}
		if sum > n {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

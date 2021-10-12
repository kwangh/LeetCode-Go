package leetcode

func guess(num int) int

func guessNumber(n int) int {
	l, r := 1, n
	m := l + (r-l)/2
	for guess(m) != 0 {
		if guess(m) > 0 {
			l = m + 1
		} else {
			r = m - 1
		}
		m = l + (r-l)/2
	}
	return m
}

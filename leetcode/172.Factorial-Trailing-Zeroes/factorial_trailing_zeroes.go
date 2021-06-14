package leetcode

func trailingZeroes(n int) int {
	var res int
	for n != 0 {
		res += n / 5
		n /= 5
	}
	return res
}

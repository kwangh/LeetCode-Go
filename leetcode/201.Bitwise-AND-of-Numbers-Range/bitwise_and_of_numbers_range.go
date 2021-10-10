package leetcode

func rangeBitwiseAnd(left int, right int) int {
	var moves int
	for left != right {
		left >>= 1
		right >>= 1
		moves++
	}
	return left << moves
}

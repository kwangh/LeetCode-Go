package leetcode

func hammingWeight(num uint32) int {
	var res int
	for ; num != 0; res++ {
		num &= num - 1
	}
	return res
}

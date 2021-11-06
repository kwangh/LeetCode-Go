package leetcode

func singleNumber(nums []int) []int {
	var xor int
	for _, v := range nums {
		xor ^= v
	}
	xor &= -xor
	res := make([]int, 2)
	for _, v := range nums {
		if xor&v == 0 {
			res[0] ^= v
		} else {
			res[1] ^= v
		}
	}
	return res
}

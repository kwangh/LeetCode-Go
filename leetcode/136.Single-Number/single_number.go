package leetcode

func singleNumber(nums []int) int {
	var res int
	for _, v := range nums {
		res ^= v
	}
	return res
}

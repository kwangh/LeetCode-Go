package leetcode

// TwoSum return index of two numbers which sums up to target val.
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	var res []int
	for i, n := range nums {
		if _, ok := m[target-n]; ok {
			res = []int{m[target-n], i}
			break
		}
		m[n] = i
	}
	return res
}

package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestConsecutive(nums []int) int {
	s := make(map[int]bool, len(nums))
	for _, v := range nums {
		s[v] = true
	}
	var res int
	for _, v := range nums {
		if !s[v-1] {
			cur := 1
			for s[v+1] {
				cur++
				v++
			}
			res = max(cur, res)
		}
	}
	return res
}

package leetcode

func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums)
	for l < r {
		m := l + (r-l)/2
		if nums[m] < target {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

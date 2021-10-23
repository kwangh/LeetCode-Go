package leetcode

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[r] {
			l = m + 1
		} else if nums[m] < nums[r] {
			r = m
		} else {
			if nums[r-1] > nums[r] {
				l = r
				break
			}
			r--
		}
	}
	return nums[l]
}

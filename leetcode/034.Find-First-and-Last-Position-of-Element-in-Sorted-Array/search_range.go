package leetcode

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if target < nums[m] {
			r = m - 1
		} else if target > nums[m] {
			l = m + 1
		} else {
			l, r = m, m
			for l >= 0 && nums[l] == target {
				l--
			}
			for r < len(nums) && nums[r] == target {
				r++
			}

			return []int{l + 1, r - 1}
		}
	}

	return []int{-1, -1}
}

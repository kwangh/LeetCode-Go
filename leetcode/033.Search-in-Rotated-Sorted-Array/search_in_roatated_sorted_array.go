package leetcode

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) / 2
		if nums[m] > nums[r] {
			l = m + 1
		} else {
			r = m
		}
	}

	s, l, r := l, 0, len(nums)-1
	if target <= nums[r] {
		l = s
	} else {
		r = s - 1
	}

	for l <= r {
		m := l + (r-l)/2
		if nums[m] == target {
			return m
		}

		if nums[m] > target {
			r = m - 1
		} else {
			l = m + 1
		}
	}

	return -1
}

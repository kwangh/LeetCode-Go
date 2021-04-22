package leetcode

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	max, curMax, curMin := nums[0], nums[0], nums[0]
	for _, v := range nums[1:] {
		if v < 0 {
			curMax, curMin = curMin, curMax
		}
		curMax = Max(v, curMax*v)
		curMin = Min(v, curMin*v)

		if curMax > max {
			max = curMax
		}
	}
	return max
}

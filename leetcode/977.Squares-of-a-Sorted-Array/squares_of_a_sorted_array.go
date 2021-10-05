func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	res := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[l]) > abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}
	return res
}
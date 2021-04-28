package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	cur, prev := 0, nums[0]
	for i := 2; i <= len(nums); i++ {
		cur = max(prev, cur+nums[i-1])
		prev, cur = cur, prev
	}
	return prev
}

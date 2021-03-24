package leetcode

func maxSubArray(nums []int) int {
	max, pre := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if pre+nums[i] < nums[i] {
			pre = nums[i]
		} else {
			pre = pre + nums[i]
		}
		if pre > max {
			max = pre
		}
	}
	return max
}

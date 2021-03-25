package leetcode

func canJump(nums []int) bool {
	lastIndex := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if i+nums[i] >= lastIndex {
			lastIndex = i
		}
	}

	return lastIndex == 0
}

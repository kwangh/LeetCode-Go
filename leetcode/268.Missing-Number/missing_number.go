package leetcode

func missingNumber(nums []int) int {
	sumFull := len(nums) * (len(nums) + 1) / 2
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	return sumFull - sum
}

package leetcode

// TwoSum return index of two numbers which sums up to target val.
func TwoSum(nums []int, target int) []int {
	var a []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return a
}

package leetcode

func removeDuplicates(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[i-res] = nums[i]
		} else {
			res++
		}
	}
	return len(nums) - res
}

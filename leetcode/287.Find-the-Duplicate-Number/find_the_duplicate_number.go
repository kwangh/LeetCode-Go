package leetcode

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i+1 {
			if nums[i] == nums[nums[i]-1] {
				return nums[i]
			}
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	return nums[len(nums)-1]
}

func findDuplicate2(nums []int) int {
	c := make([]int, len(nums))
	var res int
	for _, v := range nums {
		if c[v] > 0 {
			res = v
			break
		}
		c[v]++
	}
	return res
}

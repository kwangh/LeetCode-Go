package leetcode

func moveZeroes(nums []int) {
	var zero int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if nums[zero] != 0 {
				zero = i
			}
		} else {
			if nums[zero] == 0 {
				nums[zero], nums[i] = nums[i], nums[zero]
				zero++
			}
		}
	}
}

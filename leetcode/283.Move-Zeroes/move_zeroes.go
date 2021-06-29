package leetcode

func moveZeroes(nums []int) {
	var lastNonZero int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero] = nums[i]
			lastNonZero++
		}
	}

	for i := lastNonZero; i < len(nums); i++ {
		nums[i] = 0
	}
}

func swapZeroes(nums []int) {
	var lastNonZero int
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastNonZero], nums[i] = nums[i], nums[lastNonZero]
			lastNonZero++
		}
	}
}

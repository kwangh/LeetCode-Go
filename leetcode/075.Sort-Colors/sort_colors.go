package leetcode

func sortColors(nums []int) {
	z, o, t := 0, 0, len(nums)-1
	for o <= t {
		if nums[o] == 0 {
			nums[z], nums[o] = nums[o], nums[z]
			z++
			o++
		} else if nums[o] == 1 {
			o++
		} else {
			nums[o], nums[t] = nums[t], nums[o]
			t--
		}
	}

	for i := range nums {
		for i < t && nums[i] == 2 {
			nums[t], nums[i] = nums[i], nums[t]
			t--
		}
		if nums[i] == 0 {
			if i > z {
				nums[z], nums[i] = nums[i], nums[z]
			}
			z++
		}
	}
}

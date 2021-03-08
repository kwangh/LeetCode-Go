package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	if len(nums) < 3 {
		return res
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r, sum := i+1, len(nums)-1, -nums[i]
		for l < r {
			if nums[l]+nums[r] == sum {
				res = append(res, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			} else if nums[l]+nums[r] < sum {
				l++
			} else {
				r--
			}
		}
	}

	return res
}

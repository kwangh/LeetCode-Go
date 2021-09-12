package leetcode

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {
	var res [][]int
	if len(nums) < 4 {
		return res
	}
	sort.Ints(nums)
	for a := 0; a < len(nums)-3; a++ {
		if nums[a]+nums[a+1]+nums[a+2]+nums[a+3] > target {
			break
		}

		if nums[a]+nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] < target {
			continue
		}

		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		for b := a + 1; b < len(nums)-2; b++ {
			if nums[a]+nums[b]+nums[b+1]+nums[b+2] > target {
				break
			}

			if nums[a]+nums[b]+nums[len(nums)-1]+nums[len(nums)-2] < target {
				continue
			}

			if b > a+1 && nums[b] == nums[b-1] {
				continue
			}

			c, d := b+1, len(nums)-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]

				if sum < target {
					c++
				} else if sum == target {
					res = append(res, []int{nums[a], nums[b], nums[c], nums[d]})
					c++
					d--
					for c < d && nums[c] == nums[c-1] {
						c++
					}
					for c < d && nums[d] == nums[d+1] {
						d--
					}
				} else {
					d--
				}
			}
		}
	}

	return res
}

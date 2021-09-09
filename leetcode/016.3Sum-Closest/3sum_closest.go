package leetcode

import (
	"sort"
)

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := nums[0] + nums[1] + nums[2]
	var a, b, c int
	for a < len(nums)-2 {
		b, c = a+1, len(nums)-1
		for b < c {
			sum := nums[a] + nums[b] + nums[c]
			if abs(target-sum) < abs(target-res) {
				res = sum
			}
			if sum < target {
				b++
			} else if sum == target {
				return target
			} else {
				c--
			}
		}
		a++
	}
	return res
}

package leetcode

import (
	"sort"
)

func recursive(nums, cur []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i, v := range nums {
		if i > 0 && v == nums[i-1] {
			continue
		}
		recursive(append(append([]int{}, nums[:i]...), nums[i+1:]...), append(cur, v), res)
	}
}

func permuteUnique(nums []int) [][]int {
	var res [][]int
	sort.Ints(nums)
	recursive(nums, nil, &res)
	return res
}

package leetcode

import (
	"sort"
)

func recursive(candidates []int, target int, res *[][]int, cur []int) {
	if target == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i, v := range candidates {
		if i > 0 && v == candidates[i-1] {
			continue
		}
		if v <= target {
			recursive(candidates[i+1:], target-v, res, append(cur, v))
		} else {
			break
		}
	}
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var res [][]int
	recursive(candidates, target, &res, nil)
	return res
}

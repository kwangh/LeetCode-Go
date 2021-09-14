package combinationsum

import (
	"sort"
)

func recursiveCombinations(candidates []int, target int, res *[][]int, cur []int) {
	if target == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}

	for i, v := range candidates {
		if v <= target {
			recursiveCombinations(candidates[i:], target-v, res, append(cur, v))
		}
	}
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	sort.Sort(sort.Reverse(sort.IntSlice(candidates)))
	recursiveCombinations(candidates, target, &res, nil)
	return res
}

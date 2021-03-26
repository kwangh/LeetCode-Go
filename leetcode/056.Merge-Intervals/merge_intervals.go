package leetcode

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(a, b int) bool {
		return intervals[a][0] < intervals[b][0]
	})
	res := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if res[len(res)-1][1] >= intervals[i][0] {
			if res[len(res)-1][1] < intervals[i][1] {
				res[len(res)-1] = []int{res[len(res)-1][0], intervals[i][1]}
			}
		} else {
			res = append(res, intervals[i])
		}
	}
	return res
}

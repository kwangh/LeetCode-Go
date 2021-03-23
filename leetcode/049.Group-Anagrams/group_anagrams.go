package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := make(map[string][]string)

	for _, str := range strs {
		r := []rune(str)
		sort.Slice(r, func(i, j int) bool {
			return r[i] < r[j]
		})
		if v, exists := m[string(r)]; !exists {
			m[string(r)] = []string{str}
		} else {
			m[string(r)] = append(v, str)
		}
	}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}

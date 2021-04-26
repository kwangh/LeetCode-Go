package leetcode

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strNums := make([]string, 0, len(nums))
	for _, v := range nums {
		strNums = append(strNums, strconv.Itoa(v))
	}

	sort.Slice(strNums, func(i, j int) bool {
		s1, s2 := strNums[i]+strNums[j], strNums[j]+strNums[i]
		return s1 > s2
	})

	if strNums[0] == "0" {
		return "0"
	}
	return strings.Join(strNums, "")
}

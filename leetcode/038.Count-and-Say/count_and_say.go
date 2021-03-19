package leetcode

import (
	"strconv"
)

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	runes := []rune(countAndSay(n - 1))
	res := ""
	c := runes[0]
	count := 0
	for _, r := range runes {
		if r == c {
			count++
		} else {
			res += strconv.Itoa(count) + string(c)
			c = r
			count = 1
		}
	}
	res += strconv.Itoa(count) + string(c)
	return string(res)
}

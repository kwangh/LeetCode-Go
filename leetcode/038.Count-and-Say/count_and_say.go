package leetcode

import (
	"fmt"
	"strconv"
)

func count(n int) string {
	if n == 1 {
		return "1"
	}

	s := count(n - 1)
	fmt.Println(s)
	r := []rune(s)
	res := ""
	c := r[0]
	count := 1
	for i := 1; i < len(r); i++ {
		if r[i] == c {
			count++
		} else {
			res += strconv.Itoa(count) + string(c)
			c = r[i]
			count = 1
		}
	}
	return string(res)
}

func countAndSay(n int) string {
	return count(n)
}

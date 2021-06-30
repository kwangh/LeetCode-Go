package leetcode

import (
	"strconv"
	"strings"
)

func fizzBuzz(n int) []string {
	res := make([]string, n)
	for i := 1; i <= n; i++ {
		var sb strings.Builder
		if i%3 == 0 {
			sb.WriteString("Fizz")
		}
		if i%5 == 0 {
			sb.WriteString("Buzz")
		}
		if sb.Len() == 0 {
			res[i-1] = strconv.Itoa(i)
		} else {
			res[i-1] = sb.String()
		}
	}
	return res
}

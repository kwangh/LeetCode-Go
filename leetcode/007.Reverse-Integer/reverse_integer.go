package leetcode

import "math"

func reverse(x int) int {
	ret := 0
	for x != 0 {
		ret = ret*10 + x%10
		x = x / 10
	}
	if ret < int(-math.Pow(2, 31)) || ret > int(math.Pow(2, 31)-1) {
		return 0
	}
	return ret
}

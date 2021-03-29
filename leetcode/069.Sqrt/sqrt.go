package leetcode

func mySqrt(x int) int {
	l, r := 1, x
	for l <= r {
		mid := l + (r-l)/2
		if mid == x/mid {
			return mid
		} else if mid < x/mid {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return r
}

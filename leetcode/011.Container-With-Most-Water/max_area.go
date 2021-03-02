package leetcode

import "fmt"

// Min minimum
func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

// Max maximum
func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxArea https://leetcode.com/problems/container-with-most-water/
func MaxArea(height []int) int {
	l, r, max := 0, len(height)-1, 0
	fmt.Println(height)
	for l < r {
		max = Max(max, (r-l)*(Min(height[l], height[r])))
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return max
}

package leetcode

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestValidParentheses(s string) int {
	var left, right, max int
	for _, v := range s {
		if v == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			max = Max(max, right*2)
		} else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}

		if left == right {
			max = Max(max, right*2)
		} else if left > right {
			left, right = 0, 0
		}
	}
	return max
}

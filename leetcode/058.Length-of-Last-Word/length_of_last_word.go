package leetcode

func lengthOfLastWord(s string) int {
	var res int
	i := len(s) - 1
	for i >= 0 {
		if s[i] == ' ' {
			i--
		} else {
			break
		}
	}
	for i >= 0 {
		if s[i] != ' ' {
			res++
			i--
		} else {
			break
		}
	}
	return res
}

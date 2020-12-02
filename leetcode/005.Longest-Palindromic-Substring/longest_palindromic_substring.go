package leetcode

func LongestPalindrome(s string) string {
	length := len(s)
	max, maxStart, maxEnd := 0, 0, 0
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			start, end := i, j
			for start < end {
				if s[start] == s[end] {
					start++
					end--
				} else {
					break
				}
			}
			if start >= end {
				if max < j-i+1 {
					maxStart, maxEnd, max = i, j, j-i+1
				}
			}
		}
	}

	return s[maxStart : maxEnd+1]
}

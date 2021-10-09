package leetcode

func lengthOfLongestSubstring(s string) int {
	var max int
	chars := make([]int, 128)
	for i := range chars {
		chars[i] = -1
	}
	start := -1
	for i := range s {
		if chars[s[i]] > start {
			start = chars[s[i]]
		}
		chars[s[i]] = i
		if i-start > max {
			max = i - start
		}
	}
	return max
}

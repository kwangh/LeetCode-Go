package leetcode

func lengthOfLongestSubstring(s string) int {
	var max, start int
	for i := range s {
		j := start
		for ; j < i; j++ {
			if s[i] == s[j] {
				start = j + 1
				break
			}
		}
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}

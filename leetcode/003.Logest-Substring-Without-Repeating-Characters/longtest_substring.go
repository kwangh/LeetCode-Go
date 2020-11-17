package leetcode

func LengthOfLongestSubstring(s string) int {
	res, length := 0, len(s)
	if length == 0 {
		return 0
	}
	var cont bool
	for i := 0; i < length; i++ {
		cont = true
		for j := i + 1; j < length; j++ {
			k := i
			for k < j {
				if s[k] == s[j] {
					cont = false
					break
				}
				k++
			}

			if !cont {
				break
			}
			if res < j-i+1 {
				res = j - i + 1
			}
		}
	}
	if res == 0 {
		return 1
	}
	return res
}

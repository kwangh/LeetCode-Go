package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var chars, chars2 [26]int
	for i := range s1 {
		chars[s1[i]-'a']++
		chars2[s2[i]-'a']++
	}
	var cnt int
	for i := 0; i < 26; i++ {
		if chars[i] == chars2[i] {
			cnt++
		}
	}

	for i := 0; i < len(s2)-len(s1); i++ {
		if cnt == 26 {
			return true
		}
		l, r := s2[i]-'a', s2[i+len(s1)]-'a'
		chars2[l]--
		if chars2[l] == chars[l] {
			cnt++
		} else if chars2[l]+1 == chars[l] {
			cnt--
		}
		chars2[r]++
		if chars2[r] == chars[r] {
			cnt++
		} else if chars2[r]-1 == chars[r] {
			cnt--
		}
	}
	return cnt == 26
}

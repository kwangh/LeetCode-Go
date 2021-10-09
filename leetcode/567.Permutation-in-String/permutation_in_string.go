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
	for l := 0; l < len(s2)-len(s1); l++ {
		if match(chars, chars2) {
			return true
		}
		chars2[s2[l+len(s1)]-'a']++
		chars2[s2[l]-'a']--
	}
	return match(chars, chars2)
}

func match(s1, s2 [26]int) bool {
	for i := 0; i < 26; i++ {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

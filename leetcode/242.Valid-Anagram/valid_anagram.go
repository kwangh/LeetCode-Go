package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make([]int, 26)
	for i, v := range s {
		m[v-'a']++
		m[t[i]-'a']--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

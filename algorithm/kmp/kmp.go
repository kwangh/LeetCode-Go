package kmp

func kmp(s, pattern string) bool {
	m := make([]int, len(pattern))
	for i, j := 0, 1; j < len(pattern); j++ {
		if s[i] == s[j] {
			i++
			m[j] = i
		} else {
			if i == 0 {
				m[j] = 0
			} else {
				i = m[i-1]
			}
		}
	}

	for i, j := 0, 0; i < len(s); {
		if s[i] == pattern[j] {
			j++
			i++
			if j == len(pattern) {
				return true
			}
		} else {
			if j == 0 {
				i++
			} else {
				j = m[j-1]
			}
		}
	}
	return false
}

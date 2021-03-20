package leetcode

func isMatch(s string, p string) bool {
	i, j, star, ii := 0, 0, -1, 0
	for i < len(s) {
		if j < len(p) && (s[i] == p[j] || p[j] == '?') {
			i++
			j++
			continue
		}
		if j < len(p) && p[j] == '*' {
			star = j
			ii = i
			j++
			continue
		}
		if star != -1 {
			ii++
			i = ii
			j = star + 1
			continue
		}
		return false
	}

	for j < len(p) && p[j] == '*' {
		j++
	}
	return j == len(p)
}

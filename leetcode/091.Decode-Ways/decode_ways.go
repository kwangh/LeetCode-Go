package leetcode

func decode(s string) bool {
	if s[0] == '0' {
		return false
	}

	var n int
	if len(s) == 2 {
		n += int(s[0]-'0') * 10
		n += int(s[1] - '0')
	} else {
		n += int(s[0] - '0')
	}

	if n >= 1 && n <= 26 {
		return true
	}

	return false
}

func numDecodings(s string) int {
	m := make([]int, len(s)+1)
	m[0] = 1
	if s[0] != '0' {
		m[1] = 1
	}
	for i := 2; i <= len(s); i++ {
		if decode(string(s[i-1])) {
			m[i] += m[i-1]
		}
		if decode(string(s[i-2 : i])) {
			m[i] += m[i-2]
		}
	}

	return m[len(s)]
}

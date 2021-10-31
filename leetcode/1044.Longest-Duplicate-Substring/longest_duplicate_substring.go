package leetcode

const (
	mod  int = 1<<31 - 1
	base int = 256
)

func longestDupSubstring(s string) string {
	l, r := 0, len(s)
	var res string
	for l <= r {
		m := l + (r-l)/2
		rabin := rabinKarp(s, m)
		if rabin == "" {
			r = m - 1
		} else {
			l, res = m+1, rabin
		}
	}
	return res
}

func rabinKarp(s string, l int) string {
	if l == 0 {
		return ""
	}
	mp := make(map[int][]int)
	var h int
	h = (h*base + ord(s[0])) % mod
	base1 := 1
	for i := 1; i < l; i++ {
		h = (h*base + ord(s[i])) % mod
		base1 = (base1 * base) % mod
	}
	mp[h] = append(mp[h], 0)

	for i := 1; i+l <= len(s); i++ {
		h = ((h-base1*ord(s[i-1])+mod)*base + ord(s[i+l-1])) % mod
		for _, j := range mp[h] {
			substring := s[i : i+l]
			if s[j:j+l] == substring {
				return substring
			}
		}
		mp[h] = append(mp[h], i)
	}
	return ""
}

func ord(b byte) int {
	return int(b - 'a')
}

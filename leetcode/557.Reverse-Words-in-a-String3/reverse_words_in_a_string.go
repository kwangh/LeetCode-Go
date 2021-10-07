package leetcode

func reverseWords(s string) string {
	var start int
	res := []byte(s)
	for i, v := range s {
		if v == ' ' {
			reverse(res[start:i])
			start = i + 1
		}
	}
	reverse(res[start:len(s)])
	return string(res)
}

func reverse(b []byte) {
	for l, r := 0, len(b)-1; l < r; l, r = l+1, r-1 {
		b[l], b[r] = b[r], b[l]
	}
}

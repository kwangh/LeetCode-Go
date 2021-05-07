package leetcode

func calculate(s string) int {
	cur, last, res := 0, 0, 0
	op := byte('+')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			cur = cur*10 + int(c-'0')
		}
		if !(c >= '0' && c <= '9') && c != ' ' || i == len(s)-1 {
			if op == '+' {
				res += last
				last = cur
			} else if op == '-' {
				res += last
				last = -cur
			} else if op == '*' {
				last *= cur
			} else if op == '/' {
				last /= cur
			}
			op = c
			cur = 0
		}
	}
	res += last
	return res
}

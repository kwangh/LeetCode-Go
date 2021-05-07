package leetcode

func calculate(s string) int {
	var stack []int
	cur := 0
	op := byte('+')
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			cur = cur*10 + int(c-'0')
		}
		if !(c >= '0' && c <= '9') && c != ' ' || i == len(s)-1 {
			if op == '+' {
				stack = append(stack, cur)
			} else if op == '-' {
				stack = append(stack, -cur)
			} else if op == '*' {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop*cur)
			} else if op == '/' {
				pop := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, pop/cur)
			}
			op = c
			cur = 0
		}
	}
	var res int
	for _, v := range stack {
		res += v
	}
	return res
}

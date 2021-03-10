package leetcode

func generate(s string, left int, right int, n int, ss *[]string) {
	if left == n && right == n {
		*ss = append(*ss, s)
		return
	}

	if left < n {
		generate(s+"(", left+1, right, n, ss)
	}
	if right < left {
		generate(s+")", left, right+1, n, ss)
	}
}

// GenerateParenthesis Given n pairs of parentheses, write a function to
// generate all combinations of well-formed parentheses.
func GenerateParenthesis(n int) []string {
	res := []string{}
	generate("", 0, 0, n, &res)

	return res
}

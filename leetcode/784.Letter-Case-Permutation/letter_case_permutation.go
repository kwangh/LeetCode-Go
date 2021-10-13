package leetcode

func letterCasePermutation(s string) []string {
	var res []string
	permute(&res, s, make([]byte, len(s)), 0)
	return res
}

func permute(res *[]string, s string, chars []byte, index int) {
	if index == len(s) {
		*res = append(*res, string(chars))
		return
	}

	if isLowerLetter(s[index]) {
		chars[index] = 'A' + (s[index] - 'a')
		permute(res, s, chars, index+1)
	} else if isUpperLetter(s[index]) {
		chars[index] = 'a' + (s[index] - 'A')
		permute(res, s, chars, index+1)
	}
	chars[index] = s[index]
	permute(res, s, chars, index+1)
}

func isLowerLetter(b byte) bool {
	if b >= 'a' && b <= 'z' {
		return true
	}
	return false
}

func isUpperLetter(b byte) bool {
	if b >= 'A' && b <= 'Z' {
		return true
	}
	return false
}

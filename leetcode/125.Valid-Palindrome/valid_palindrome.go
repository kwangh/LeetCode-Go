package leetcode

func isAlphanumeric(b byte) bool {
	if (b >= '0' && b <= '9') || (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
		return true
	}
	return false
}

func lowercase(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		for i < len(s)-1 && !isAlphanumeric(s[i]) {
			i++
		}
		for j > 0 && !isAlphanumeric(s[j]) {
			j--
		}
		if i < j && lowercase(s[i]) != lowercase(s[j]) {
			return false
		}
		i++
		j--
	}
	return true
}

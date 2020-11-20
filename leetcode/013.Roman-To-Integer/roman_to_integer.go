package leetcode

// RomanCharToInt changes byte to int
func RomanCharToInt(b byte) int {
	switch b {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}

// RomanToInt changes roman letters to integer value
func RomanToInt(s string) int {
	res := 0
	length := len(s)
	for i := 0; i < length; i++ {
		if i+1 < length && RomanCharToInt(s[i]) < RomanCharToInt(s[i+1]) {
			res += (RomanCharToInt(s[i+1]) - RomanCharToInt(s[i]))
			i++
		} else {
			res += RomanCharToInt(s[i])
		}
	}
	return res
}

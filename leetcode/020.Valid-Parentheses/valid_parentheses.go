package leetcode

// IsValid return true/false based on right parentheses rule
func IsValid(s string) bool {
	stack := make([]rune, len(s))
	length := 0
	for _, r := range s {
		if length == 0 {
			if r == ')' || r == '}' || r == ']' {
				return false
			}
			stack[length] = r
			length++
			continue
		}

		if r == ')' {
			if stack[length-1] == '(' {
				length--
			} else {
				return false
			}
		} else if r == '}' {
			if stack[length-1] == '{' {
				length--
			} else {
				return false
			}
		} else if r == ']' {
			if stack[length-1] == '[' {
				length--
			} else {
				return false
			}
		} else {
			stack[length] = r
			length++
		}
	}

	if length == 0 {
		return true
	}
	return false
}

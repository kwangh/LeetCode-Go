package leetcode

import (
	"fmt"
	"strconv"
)

//IsDigit check byte if it is digit
func IsDigit(b byte) bool {
	if b >= '0' && b <= '9' {
		return true
	}

	return false
}

//MyAtoi string to int
func MyAtoi(s string) int {
	ns := ""
	IsFirtDigit := false
	for i := 0; i < len(s); i++ {
		if !IsFirtDigit {
			switch s[i] {
			case ' ':
				continue
			case '+':
				IsFirtDigit = true
				continue
			case '-':
				ns = "-"
				IsFirtDigit = true
				continue
			}

			if IsDigit(s[i]) {
				ns = fmt.Sprintf("%s%c", ns, s[i])
				IsFirtDigit = true
			} else {
				return 0
			}
		} else {
			if IsDigit(s[i]) {
				ns = fmt.Sprintf("%s%c", ns, s[i])
			} else {
				break
			}
		}
	}
	i, _ := strconv.ParseInt(ns, 10, 64)
	if i < -2147483648 {
		return -2147483648
	} else if i > 2147483647 {
		return 2147483647
	}
	return int(i)
}

package leetcode

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var sb strings.Builder
	for i := 0; i < numRows; i++ {
		for j := 0; j+i < len(s); j += 2*numRows - 2 {
			sb.WriteByte(s[i+j])
			if i != 0 && i != numRows-1 && j+2*numRows-2-i < len(s) {
				sb.WriteByte(s[j+2*numRows-2-i])
			}
		}
	}
	return sb.String()
}

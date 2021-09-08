package leetcode

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	m := make([][]rune, numRows)
	for i := range m {
		m[i] = make([]rune, 1000)
	}
	x, y, down := 0, 0, true
	for _, r := range s {
		m[y][x] = r
		if down {
			if y+1 < numRows {
				y++
			} else {
				down = false
				x++
				y--
			}
		} else {
			if y-1 >= 0 {
				y--
				x++
			} else {
				down = true
				y++
			}
		}
	}

	var sb strings.Builder
	for i := range m {
		for j := range m[i] {
			if m[i][j] != 0 {
				sb.WriteRune(m[i][j])
			}
		}
	}

	return sb.String()
}

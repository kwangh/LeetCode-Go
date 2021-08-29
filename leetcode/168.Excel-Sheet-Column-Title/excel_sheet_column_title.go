package leetcode

func convertToTitle(columnNumber int) string {
	var runes []rune
	for columnNumber > 0 {
		columnNumber--
		runes = append([]rune{rune('A' + (columnNumber)%26)}, runes...)
		columnNumber /= 26
	}
	return string(runes)
}

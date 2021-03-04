package leetcode

// LetterCombinations letter combinations
func LetterCombinations(digits string) []string {
	letters := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"},
		{"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"},
		{"t", "u", "v"}, {"w", "x", "y", "z"}}
	res := []string{}
	for _, r := range digits {
		tmp := []string{}
		for _, l := range letters[r-'2'] {
			if len(res) == 0 {
				tmp = append(tmp, l)
			} else {
				for _, c := range res {
					tmp = append(tmp, c+l)
				}
			}
		}
		res = tmp
	}
	return res
}

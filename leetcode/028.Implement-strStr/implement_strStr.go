package leetcode

func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	haystackRunes := []rune(haystack)
	needleRunes := []rune(needle)
	for i := range haystack {
		if i+len(needle) > len(haystack) {
			break
		}

		for j := 0; j < len(needle); j++ {
			if haystackRunes[i+j] != needleRunes[j] {
				break
			}
			if j == len(needle)-1 {
				return i
			}
		}
	}

	return -1
}

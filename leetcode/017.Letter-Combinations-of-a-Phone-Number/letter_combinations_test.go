package leetcode

import (
	"sort"
	"testing"
)

func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// LetterCombinations letter combinations
func TestLetterCombinations(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{"23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
		{"", []string{}},
		{"2", []string{"a", "b", "c"}},
	}

	for _, c := range cases {
		got := LetterCombinations(c.in)
		sort.Strings(got)
		if !equal(got, c.want) {
			t.Errorf("in: %s, \texpected: %v, instead got: %v\n", c.in, c.want, got)
		}
	}
}

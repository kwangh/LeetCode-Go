package leetcode

import "testing"

func TestLongestPalindrome(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}
	for _, c := range cases {
		got := LongestPalindrome(c.in)
		if got != c.want {
			t.Errorf("expected %s instead got %s\n", c.want, got)
		}
	}
}

package kmp

import "testing"

func TestKmp(t *testing.T) {
	cases := []struct {
		s, pattern string
		want       bool
	}{
		{"abxabcabcaby", "abcaby", true},
		{"abcbcglx", "bcgl", true},
	}
	for _, c := range cases {
		got := kmp(c.s, c.pattern)
		if got != c.want {
			t.Errorf("want:%t instead got:%t", c.want, got)
		}
	}
}

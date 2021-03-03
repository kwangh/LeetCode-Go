package leetcode

import "testing"

// IsMatch https://leetcode.com/problems/regular-expression-matching/
func TestIsMatch(t *testing.T) {
	cases := []struct {
		in      string
		pattern string
		want    bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"ab", ".*c", false},
		{"aaa", "a*a", true},
	}

	for _, c := range cases {
		got := IsMatch(c.in, c.pattern)
		if got != c.want {
			t.Errorf("in, pattern: %s, %s,\texpected: %t, instead got: %t\n", c.in, c.pattern, c.want, got)
		}
	}
}

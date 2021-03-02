package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
		{"{[[(({{", false},
	}

	for _, c := range cases {
		got := IsValid(c.in)
		if got != c.want {
			t.Errorf("input: %s\texpected %t instead got %t\n", c.in, c.want, got)
		}
	}
}

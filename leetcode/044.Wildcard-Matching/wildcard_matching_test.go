package leetcode

import "testing"

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s    string
		p    string
		want bool
	}{

		{
			s:    "abcdefdef",
			p:    "abc*def",
			want: true,
		},
	}
	for _, c := range cases {
		got := isMatch(c.s, c.p)
		if c.want != got {
			t.Errorf("s:%s\tp:%s\twant:%t instead got:%t", c.s, c.p, c.want, got)
		}
	}
}

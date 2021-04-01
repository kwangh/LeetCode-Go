package leetcode

import "testing"

func TestNumDecodings(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{
			"12", 2,
		},
		{
			"226", 3,
		},
		{
			"10", 1,
		},
		{
			"27", 1,
		},
		{
			"06", 0,
		},
		{
			"13027", 0,
		},
	}

	for _, c := range cases {
		got := numDecodings(c.s)
		if c.want != got {
			t.Errorf("string:%s\twant:%d instead got %d", c.s, c.want, got)
		}
	}
}

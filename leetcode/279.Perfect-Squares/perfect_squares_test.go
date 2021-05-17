package leetcode

import "testing"

func TestNumSquares(t *testing.T) {
	cases := []struct {
		n, want int
	}{
		{12, 3},
		{13, 2},
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 1},
	}
	for _, c := range cases {
		got := numSquares(c.n)
		if c.want != got {
			t.Errorf("n:%d\twant:%d instead got:%d", c.n, c.want, got)
		}
	}
}

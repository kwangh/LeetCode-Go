package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	cases := []struct {
		m, n int
		want int
	}{
		{
			3, 7, 28,
		},
		{
			3, 2, 3,
		},
		{
			7, 3, 28,
		},
		{
			3, 3, 6,
		},
	}
	for _, c := range cases {
		got := uniquePaths(c.m, c.n)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	cases := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
	}
	for _, c := range cases {
		got := climbStairs(c.n)
		if got != c.want {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

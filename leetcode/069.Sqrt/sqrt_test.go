package leetcode

import "testing"

func TestMySqrt(t *testing.T) {
	cases := []struct {
		x    int
		want int
	}{
		{4, 2}, {8, 2}, {1, 1}, {2, 1},
	}
	for _, c := range cases {
		got := mySqrt(c.x)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

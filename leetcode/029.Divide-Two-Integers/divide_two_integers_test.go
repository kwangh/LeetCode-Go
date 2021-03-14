package leetcode

import "testing"

func TestDivide(t *testing.T) {
	cases := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{
			10, 3, 3,
		},
		{
			7, -3, -2,
		},
		{
			0, 1, 0,
		},
		{
			1, 1, 1,
		},
		{
			MinInt32, -1, MaxInt32,
		},
		{
			MinInt32, MinInt32, 1,
		},
	}
	for _, c := range cases {
		got := divide(c.dividend, c.divisor)
		t.Logf("dividend:%d, divisor:%d, want:%d, got:%d\n",c.dividend,c.divisor,c.want,got)
		if got != c.want {
			t.Errorf("want:%d, instead got:%d\n", c.want, got)
		}
	}
}

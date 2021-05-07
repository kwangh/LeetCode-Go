package leetcode

import "testing"

func TestCalculate(t *testing.T) {
	cases := []struct {
		s    string
		want int
	}{
		{"3+2*2", 7},
		{" 3/2 ", 1},
		{" 3+5 / 2 ", 5},
	}
	for _, c := range cases {
		got := calculate(c.s)
		if c.want != got {
			t.Errorf("%s=%d(Correct answer:%d)", c.s, got, c.want)
		}
	}
}

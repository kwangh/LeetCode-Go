package leetcode

import "testing"

func TestMaxPoints(t *testing.T) {
	cases := []struct {
		in   [][]int
		want int
	}{
		{
			in:   [][]int{{1, 1}, {2, 2}, {3, 3}},
			want: 3,
		},
		{
			in:   [][]int{{1, 1}, {3, 2}, {5, 3}, {4, 1}, {2, 3}, {1, 4}},
			want: 4,
		},
	}
	for _, c := range cases {
		got := maxPoints(c.in)
		if c.want != got {
			t.Errorf("want:%d instead got:%d\n", c.want, got)
		}
	}
}

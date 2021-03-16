package leetcode

import "testing"

func TestTrap(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{
			in:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			want: 6,
		},
		{
			in:   []int{4, 2, 0, 3, 2, 5},
			want: 9,
		},
		{
			in:   []int{1, 1, 1, 1, 1},
			want: 0,
		},
		{
			in:   []int{1, 2, 3, 4, 5},
			want: 0,
		},
		{
			in:   []int{5, 4, 3, 2, 1},
			want: 0,
		},
		{
			in:   []int{1000, 0, 100},
			want: 100,
		},
	}
	for _, c := range cases {
		got := trap(c.in)
		if c.want != got {
			t.Errorf("want:%d, instead got:%d\n", c.want, got)
		}
	}
}

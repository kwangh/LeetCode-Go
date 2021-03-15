package leetcode

import "testing"

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{
			[]int{1, 2, 0}, 3,
		},
		{
			[]int{3, 4, -1, 1}, 2,
		},
		{
			[]int{7, 8, 9, 11, 12}, 1,
		},
	}
	for _, c := range cases {
		got := firstMissingPositive(c.in)
		if got != c.want {
			t.Errorf("want:%d, instead got:%d\n", c.want, got)
		}
	}
}

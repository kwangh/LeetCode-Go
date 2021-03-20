package leetcode

import "testing"

func TestPermute(t *testing.T) {
	cases := []struct {
		nums []int
		want [][]int
	}{
		{
			[]int{1, 2, 3},
			[][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}},
		},
		{
			[]int{0, 1},
			[][]int{{0, 1}, {1, 0}},
		},
		{
			[]int{1},
			[][]int{{1}},
		},
	}
	for _, c := range cases {
		got := permute(c.nums)
		if len(got) != len(c.want) {
			t.Errorf("want:%v instead got:%v", c.want, got)
		}
	}
}

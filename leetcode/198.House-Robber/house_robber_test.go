package leetcode

import "testing"

func TestRob(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{[]int{1}, 1},
		{[]int{400, 0, 1, 400}, 800},
	}
	for _, c := range cases {
		got := rob(c.nums)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

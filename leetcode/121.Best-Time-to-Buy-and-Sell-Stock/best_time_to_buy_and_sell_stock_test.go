package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}
	for _, c := range cases {
		got := maxProfit(c.prices)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

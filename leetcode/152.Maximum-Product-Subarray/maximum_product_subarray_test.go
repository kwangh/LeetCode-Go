package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}
	for _, c := range cases {
		got := maxProduct(c.nums)
		if c.want != got {
			t.Errorf("nums:%v\twant:%d instead got:%d", c.nums, c.want, got)
		}
	}
}

package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{
			[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			6,
		},
		{
			[]int{1},
			1,
		},
		{
			[]int{-1},
			-1,
		},
		{
			[]int{-1, 0, -1},
			0,
		},
	}
	for _, c := range cases {
		got := maxSubArray(c.nums)
		if got != c.want {
			t.Errorf("want:%d instead got:%d", c.want, got)
			break
		}
	}
}

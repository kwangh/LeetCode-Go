package leetcode

import "testing"

func TestCanJump(t *testing.T) {
	cases := []struct {
		nums []int
		want bool
	}{
		{
			[]int{2, 3, 1, 1, 4},
			true,
		},
		{
			[]int{3, 2, 1, 0, 4},
			false,
		},
		{
			[]int{1, 1, 1, 0},
			true,
		},
	}
	for _, c := range cases {
		got := canJump(c.nums)
		if c.want != got {
			t.Errorf("nums:%v\twant:%t instead got:%t", c.nums, c.want, got)
		}
	}
}

package leetcode

import "testing"

func TestFindDuplicate(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
		{[]int{1, 1}, 1},
		{[]int{1, 1, 2}, 1},
	}
	for _, c := range cases {
		got := findDuplicate(c.nums)
		if c.want != got {
			t.Errorf("nums:%v\twant:%d instead got:%d", c.nums, c.want, got)
		}
	}
}

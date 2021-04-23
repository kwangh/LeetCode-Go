package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 2},
	}
	for _, c := range cases {
		got := findPeakElement(c.nums)
		if c.want != got {
			t.Errorf("nums:%v\twant:%d instead got:%d", c.nums, c.want, got)
		}
	}
}

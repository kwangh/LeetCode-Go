package leetcode

import "testing"

func TestFindKthLargest(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, c := range cases {
		got := findKthLargest(c.nums, c.k)
		if c.want != got {
			t.Errorf("nums:%v, k:%d\twant:%d instead got:%d", c.nums, c.k, c.want, got)
		}
	}
}

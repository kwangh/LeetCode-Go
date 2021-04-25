package leetcode

import "testing"

func TestMajorityElement(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, c := range cases {
		got := majorityElement(c.nums)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

package leetcode

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func TestSortColors(t *testing.T) {
	cases := []struct {
		nums []int
		want []int
	}{
		{[]int{2, 2, 2, 2}, []int{2, 2, 2, 2}},
		{[]int{0, 1, 2, 0, 1, 2, 0, 1, 2}, []int{0, 0, 0, 1, 1, 1, 2, 2, 2}},
	}
	for _, c := range cases {
		sortColors(c.nums)
		if !equal(c.nums, c.want) {
			t.Errorf("want:%v instead got:%v", c.want, c.nums)
		}
	}
}

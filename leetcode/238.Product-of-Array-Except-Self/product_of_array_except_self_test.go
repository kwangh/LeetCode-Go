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

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, c := range cases {
		got := productExceptSelf(c.nums)
		if !equal(c.want, got) {
			t.Errorf("want:%v instead got:%v", c.want, got)
		}
	}
}

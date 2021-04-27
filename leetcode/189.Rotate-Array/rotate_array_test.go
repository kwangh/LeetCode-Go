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

func TestRotate(t *testing.T) {
	cases := []struct {
		nums, want []int
		k          int
	}{
		{nums: []int{1, 2, 3, 4, 5, 6, 7}, want: []int{5, 6, 7, 1, 2, 3, 4}, k: 3},
		{nums: []int{-1, -100, 3, 99}, want: []int{3, 99, -1, -100}, k: 2},
		{nums: []int{-1}, want: []int{-1}, k: 2},
	}
	for _, c := range cases {
		rotate(c.nums, c.k)
		if !equal(c.nums, c.want) {
			t.Errorf("want:%v instead got:%v", c.want, c.nums)
		}
	}
}

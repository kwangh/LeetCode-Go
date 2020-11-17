package leetcode

import (
	"testing"
)

func Equal(a, b []int) bool {
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

func TestTwoSum(t *testing.T) {
	cases := []struct {
		target     int
		nums, want []int
	}{
		{9, []int{2, 7, 11, 15}, []int{0, 1}},
		{6, []int{3, 2, 4}, []int{1, 2}},
	}
	for _, c := range cases {
		got := TwoSum(c.nums, c.target)
		if !Equal(c.want, got) {
			t.Error()
		}
	}
}

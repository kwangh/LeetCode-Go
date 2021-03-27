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

func TestPlusOne(t *testing.T) {
	cases := []struct {
		digits []int
		want   []int
	}{
		{
			[]int{1, 2, 3},
			[]int{1, 2, 4},
		},
		{
			[]int{4, 3, 2, 1},
			[]int{4, 3, 2, 2},
		},
		{
			[]int{0},
			[]int{1},
		},
		{
			[]int{9, 9, 9},
			[]int{1, 0, 0, 0},
		},
		{
			[]int{9},
			[]int{1, 0},
		},
	}
	for _, c := range cases {
		got := plusOne(c.digits)
		if !equal(got, c.want) {
			t.Errorf("want:%v instead got:%v", c.want, got)
		}
	}
}

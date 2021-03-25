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

func TestSpiralOrder(t *testing.T) {
	cases := []struct {
		matrix [][]int
		want   []int
	}{
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[]int{1, 2, 3, 6, 9, 8, 7, 4, 5},
		},
		{
			[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}},
			[]int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7},
		},
	}
	for _, c := range cases {
		got := spiralOrder(c.matrix)
		if !equal(got, c.want) {
			t.Errorf("want:%v instead got:%v", c.want, got)
		}
	}
}

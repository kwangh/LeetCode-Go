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

func TestSetZeroes(t *testing.T) {
	cases := []struct {
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}},
		},
		{
			[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}},
		},
	}
	for _, c := range cases {
		setZeroes(c.matrix)
		for i, row := range c.want {
			if !equal(row, c.matrix[i]) {
				t.Errorf("want:%v instead got:%v", c.want, c.matrix)
			}
		}
	}
}

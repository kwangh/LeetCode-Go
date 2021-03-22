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
		matrix [][]int
		want   [][]int
	}{
		{
			[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}},
		},
		{
			[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			[][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}},
		},
		{
			[][]int{{1}},
			[][]int{{1}},
		},
		{
			[][]int{{1, 2}, {3, 4}},
			[][]int{{3, 1}, {4, 2}},
		},
	}
	for _, c := range cases {
		rotate(c.matrix)
		for i, s := range c.matrix {
			if !equal(s, c.want[i]) {
				t.Errorf("want:%v instead got:%v", c.want, c.matrix)
				break
			}
		}
	}
}

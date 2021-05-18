package leetcode

import "testing"

func equal(a, b [][]int) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}

func TestGameOfLife(t *testing.T) {
	cases := []struct {
		board, want [][]int
	}{
		{[][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}},
		{[][]int{{1, 1}, {1, 0}}, [][]int{{1, 1}, {1, 1}}},
	}
	for _, c := range cases {
		gameOfLife(c.board)
		if !equal(c.board, c.want) {
			t.Errorf("want:%v\tinstead got:%v", c.want, c.board)
		}
	}
}

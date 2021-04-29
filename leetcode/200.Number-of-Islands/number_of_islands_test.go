package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	cases := []struct {
		grid [][]byte
		want int
	}{
		{[][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}, 1},
		{[][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}, 3},
		{[][]byte{{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}}, 1},
	}
	for _, c := range cases {
		got := numIslands(c.grid)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

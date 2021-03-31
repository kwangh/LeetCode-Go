package leetcode

import "testing"

func TestExist(t *testing.T) {
	cases := []struct {
		board [][]byte
		word  string
		want  bool
	}{
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCB",
			false,
		},
		{
			[][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}},
			"AAB",
			true,
		},
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABCESEEEFS",
			true,
		},
	}
	for _, c := range cases {
		got := exist(c.board, c.word)
		if c.want != got {
			t.Errorf("want:%t instead got:%t", c.want, got)
		}
	}
}

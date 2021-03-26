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

func TestMerge(t *testing.T) {
	cases := []struct {
		intervals [][]int
		want      [][]int
	}{
		{
			[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			[][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			[][]int{{1, 4}, {4, 5}},
			[][]int{{1, 5}},
		},
		{
			[][]int{{1, 4}, {2, 3}},
			[][]int{{1, 4}},
		},
		{
			[][]int{{1, 4}, {0, 2}, {3, 5}},
			[][]int{{0, 5}},
		},
	}
	for _, c := range cases {
		got := merge(c.intervals)
		if len(got) != len(c.want) {
			t.Errorf("intervals:%v\twant:%v instead got:%v", c.intervals, c.want, got)
			continue
		}
		for i, v := range got {
			if !equal(v, c.want[i]) {
				t.Errorf("intervals:%v\twant:%v instead got:%v", c.intervals, c.want, got)
			}
		}
	}
}

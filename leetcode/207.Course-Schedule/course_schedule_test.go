package leetcode

import "testing"

func TestCanFinish(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		want          bool
	}{
		{2, [][]int{{1, 0}}, true},
		{2, [][]int{{1, 0}, {0, 1}}, false},
		{4, [][]int{{1, 0}, {0, 2}, {2, 0}, {2, 3}}, false},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true},
	}
	for _, c := range cases {
		got := canFinish(c.numCourses, c.prerequisites)
		if c.want != got {
			t.Errorf("prerequisiteds:%v\twant:%t instead got:%t", c.prerequisites, c.want, got)
		}
	}
}

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

func TestFindOrder(t *testing.T) {
	cases := []struct {
		numCourses    int
		prerequisites [][]int
		want          []int
	}{
		{2, [][]int{{1, 0}}, []int{0, 1}},
		{4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, []int{0, 1, 2, 3}},
		{1, nil, []int{0}},
	}
	for _, c := range cases {
		got := findOrder(c.numCourses, c.prerequisites)
		if !equal(c.want, got) {
			t.Errorf("want:%v instead got:%v", c.want, got)
		}
	}
}

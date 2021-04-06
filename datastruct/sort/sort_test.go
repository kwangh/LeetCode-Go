package sort

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

func TestQuickSort(t *testing.T) {
	cases := []struct {
		s    []int
		want []int
	}{
		{s: []int{7, 1, 2, 5, 2, 6, 9, 10, 31, 21, 58, 6, 2}, want: []int{1, 2, 2, 2, 5, 6, 6, 7, 9, 10, 21, 31, 58}},
		{s: []int{1, 1, 1}, want: []int{1, 1, 1}},
		{s: []int{1, 2, 3}, want: []int{1, 2, 3}},
	}
	for _, c := range cases {
		QuickSort(c.s, 0, len(c.s)-1)
		if !equal(c.s, c.want) {
			t.Errorf("want:%v instead got:%v", c.want, c.s)
		}
	}
}

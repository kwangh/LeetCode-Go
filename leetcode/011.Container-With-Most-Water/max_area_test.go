package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
		{[]int{2, 3, 4, 5, 18, 17, 6}, 17},
		{[]int{6, 17, 18, 5, 4, 3, 2}, 17},
	}

	for _, c := range cases {
		got := MaxArea(c.in)
		if got != c.want {
			t.Errorf("expected %d instead got %d\n", c.want, got)
		}
	}
}

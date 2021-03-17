package leetcode

import "testing"

func TestSearch(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   int
	}{
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
	}
	for _, c := range cases {
		got := search(c.nums, c.target)
		if c.want != got {
			t.Errorf("want:%d instead got:%d\n", c.want, got)
		}
	}
}

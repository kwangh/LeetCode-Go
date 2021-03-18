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

func TestSearchRange(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		want   []int
	}{
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 8,
			want:   []int{3, 4},
		},
		{
			nums:   []int{5, 7, 7, 8, 8, 10},
			target: 6,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{1, 1, 1, 1},
			target: 0,
			want:   []int{-1, -1},
		},
		{
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		{
			nums:   []int{1},
			target: 1,
			want:   []int{0, 0},
		},
		{
			nums:   []int{1, 4},
			target: 4,
			want:   []int{1, 1},
		},
	}
	for _, c := range cases {
		got := searchRange(c.nums, c.target)
		if !equal(c.want, got) {
			t.Errorf("want:%d instead got:%d\n", c.want, got)
		}
	}
}

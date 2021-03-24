package leetcode

import "testing"

func TestMissingNumber(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{
			[]int{3, 0, 1},
			2,
		},
		{
			[]int{0, 1},
			2,
		},
		{
			[]int{9, 6, 4, 2, 3, 5, 7, 0, 1},
			8,
		},
		{
			[]int{0},
			1,
		},
		{
			[]int{1, 2},
			0,
		},
	}
	for _, c := range cases {
		got := missingNumber(c.nums)
		if got != c.want {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

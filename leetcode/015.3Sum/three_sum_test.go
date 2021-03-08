package leetcode

import (
	"testing"
)

func TestThreeSum(t *testing.T) {
	cases := []struct {
		in   []int
		want [][]int
	}{
		{
			in:   []int{-1, 0, 1, 2, -1, -4},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			in:   []int{},
			want: [][]int{},
		},
		{
			in:   []int{0},
			want: [][]int{},
		},
		{
			in:   []int{0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			in:   []int{0, 0, 0, 0, 0},
			want: [][]int{{0, 0, 0}},
		},
		{
			in:   []int{-2, 0, 1, 1, 2},
			want: [][]int{{-2, 0, 2}, {-2, 1, 1}},
		},
	}
	for _, c := range cases {
		got := threeSum(c.in)
		if len(got) != len(c.want) {
			t.Errorf("len(got)!=len(c.want), len(got)=%d, len(c.want)=%d", len(got), len(c.want))
		}
		t.Logf("got:%+v, c.want:%+v", got, c.want)
	}
}

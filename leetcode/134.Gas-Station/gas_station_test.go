package leetcode

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	cases := []struct {
		gas  []int
		cost []int
		want int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{2}, []int{3}, -1},
		{[]int{2}, []int{1}, 0},
	}
	for _, c := range cases {
		got := canCompleteCircuit(c.gas, c.cost)
		if got != c.want {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

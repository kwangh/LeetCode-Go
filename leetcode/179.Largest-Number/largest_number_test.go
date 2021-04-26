package leetcode

import "testing"

func TestLargestNumber(t *testing.T) {
	cases := []struct {
		nums []int
		want string
	}{
		{[]int{10, 2}, "210"},
		{[]int{3, 30, 34, 5, 9}, "9534330"},
		{[]int{1}, "1"},
		{[]int{0, 0}, "0"},
		{[]int{10, 100}, "10100"},
		{[]int{34323, 3432}, "343234323"},
		{[]int{0, 9, 8, 7, 6, 5, 4, 3, 2, 1}, "9876543210"},
	}
	for _, c := range cases {
		got := largestNumber(c.nums)
		if c.want != got {
			t.Errorf("want:%s instead got:%s", c.want, got)
		}
	}
}

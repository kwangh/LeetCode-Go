package leetcode

import (
	"testing"

	. "github.com/kwangh/leetcode/structures"
)

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

func TestMiddleNode(t *testing.T) {
	cases := []struct {
		head, want *ListNode
	}{
		{head: Lists([]int{1, 2, 3, 4, 5}), want: Lists([]int{3, 4, 5})},
		{head: Lists([]int{1, 2, 3, 4, 5, 6}), want: Lists([]int{4, 5, 6})},
	}
	for _, c := range cases {
		got := middleNode(c.head)
		gotInts := Ints(got)
		want := Ints(c.want)
		if !equal(gotInts, want) {
			t.Errorf("want:%v instead got:%v", want, gotInts)
		}
	}
}

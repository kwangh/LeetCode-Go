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

func TestSortList(t *testing.T) {
	cases := []struct {
		head, want *ListNode
	}{
		{head: Lists([]int{4, 2, 1, 3}), want: Lists([]int{1, 2, 3, 4})},
		{head: Lists([]int{-1, 5, 3, 4, 0}), want: Lists([]int{-1, 0, 3, 4, 5})},
		{head: nil, want: nil},
	}
	for _, c := range cases {
		got := sortList(c.head)
		gotInts := Ints(got)
		want := Ints(c.want)
		if !equal(gotInts, want) {
			t.Errorf("want:%v instead got:%v", want, gotInts)
		}
	}
}

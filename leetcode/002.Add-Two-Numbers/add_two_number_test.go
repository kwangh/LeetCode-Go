package leetcode

import (
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	cases := []struct {
		in1, in2, want *ListNode
	}{
		{
			&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
			&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
			&ListNode{Val: 7, Next: &ListNode{Val: 0, Next: &ListNode{Val: 8}}},
		},
	}
	for _, c := range cases {
		got := AddTwoNumbers(c.in1, c.in2)
		for c.want != nil {
			if got.Val != c.want.Val {
				t.Error()
			}
			c.want = c.want.Next
			got = got.Next
		}
	}
}

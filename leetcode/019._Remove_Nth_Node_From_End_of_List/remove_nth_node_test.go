package leetcode

import "testing"

func TestRemoveNthFromEnd(t *testing.T) {
	cases := []struct {
		in   *ListNode
		n    int
		want *ListNode
	}{
		{
			in:   &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}},
			n:    2,
			want: &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{5, nil}}}},
		},
		{
			in:   &ListNode{1, nil},
			n:    1,
			want: &ListNode{},
		},
		{
			in:   &ListNode{1, &ListNode{2, nil}},
			n:    1,
			want: &ListNode{1, nil},
		},
	}

	for _, c := range cases {
		t.Log(c.in, removeNthFromEnd(c.in, c.n))
	}
}

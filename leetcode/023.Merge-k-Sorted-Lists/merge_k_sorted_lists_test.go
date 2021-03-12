package leetcode

import "testing"

func TestMergeKLists(t *testing.T) {
	cases := []struct {
		in   []*ListNode
		want *ListNode
	}{
		{
			in: []*ListNode{
				{1, &ListNode{4, &ListNode{5, nil}}},
				{1, &ListNode{3, &ListNode{4, nil}}},
				{2, &ListNode{6, nil}},
			},
			want: &ListNode{1, &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}},
		},
		{
			in:   []*ListNode{},
			want: &ListNode{},
		},
		{
			in:   []*ListNode{{}},
			want: &ListNode{},
		},
	}

	for _, c := range cases {
		got := mergeKLists(c.in)
		if got == nil {
			t.Errorf("got is nil")
			return
		}

		for got != nil && c.want != nil {
			t.Logf("want:%d got:%d", c.want.Val, got.Val)
			if got.Val != c.want.Val {
				t.Errorf("want:%d instead got:%d", c.want.Val, got.Val)
				return
			}
			got = got.Next
			c.want = c.want.Next
		}
		if got != nil || c.want != nil {
			t.Errorf("got:%v want:%v", got, c.want)
			return
		}
	}
}

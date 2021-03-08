package leetcode

// ListNode list nodes
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	t := res
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			t.Next = &ListNode{Val: l1.Val}
			t = t.Next
			l1 = l1.Next
		} else {
			t.Next = &ListNode{Val: l2.Val}
			t = t.Next
			l2 = l2.Next
		}
	}
	for l1 != nil {
		t.Next = &ListNode{Val: l1.Val}
		t = t.Next
		l1=l1.Next
	}

	for l2 != nil {
		t.Next = &ListNode{Val: l2.Val}
		t = t.Next
		l2=l2.Next
	}
	return res.Next
}

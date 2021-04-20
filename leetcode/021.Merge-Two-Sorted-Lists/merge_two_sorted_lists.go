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
			t.Next = l1
			t = t.Next
			l1 = l1.Next
		} else {
			t.Next = l2
			t = t.Next
			l2 = l2.Next
		}
	}
	if l1 != nil {
		t.Next = l1
	} else {
		t.Next = l2
	}
	return res.Next
}

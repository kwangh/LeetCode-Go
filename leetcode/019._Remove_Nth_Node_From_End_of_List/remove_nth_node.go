package leetcode

// ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// RemoveNthFromEnd remove nth node from end
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	s := []*ListNode{}
	for head != nil {
		s = append(s, head)
		head = head.Next
	}
	if len(s) == n {
		return s[0].Next
	}
	s[len(s)-n-1].Next = s[len(s)-n].Next
	return s[0]
}

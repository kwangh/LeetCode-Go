package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	newHead := &ListNode{0, head}
	s, f := newHead, newHead
	for n > 0 {
		f = f.Next
		n--
	}

	for f.Next != nil {
		f = f.Next
		s = s.Next
	}
	s.Next = s.Next.Next
	return newHead.Next
}

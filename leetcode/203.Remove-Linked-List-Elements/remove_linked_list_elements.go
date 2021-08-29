package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	v := &ListNode{0, head}
	last := v
	for head != nil {
		if head.Val == val {
			last.Next = head.Next
		} else {
			last = last.Next
		}
		head = head.Next
	}
	return v.Next
}

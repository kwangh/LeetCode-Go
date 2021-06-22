package leetcode

import (
	"github.com/kwangh/leetcode/structures"
)

func reverseList(head *structures.ListNode) *structures.ListNode {
	var cur *structures.ListNode
	for head != nil {
		tmp := head.Next
		head.Next = cur
		cur = head
		head = tmp
	}
	return cur
}

func reverseListRecursive(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	n := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return n
}

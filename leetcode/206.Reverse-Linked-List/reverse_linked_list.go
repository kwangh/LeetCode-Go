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

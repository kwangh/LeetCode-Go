package leetcode

import (
	. "github.com/kwangh/leetcode/structures"
)

func middleNode(head *ListNode) *ListNode {
	cnt := 0
	tmp := head
	for tmp != nil {
		cnt++
		tmp = tmp.Next
	}

	for i := 0; i < cnt/2; i++ {
		head = head.Next
	}
	return head
}

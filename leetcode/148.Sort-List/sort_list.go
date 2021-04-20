package leetcode

import (
	. "github.com/kwangh/leetcode/structures"
)

func getMid(head *ListNode) *ListNode {
	var midPrev *ListNode
	for head != nil && head.Next != nil {
		if midPrev == nil {
			midPrev = head
		} else {
			midPrev = midPrev.Next
		}
		head = head.Next.Next
	}
	mid := midPrev.Next
	midPrev.Next = nil
	return mid
}

func merge(l1, l2 *ListNode) *ListNode {
	tmp := &ListNode{}
	cur := tmp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cur.Next = l1
			l1 = l1.Next
			cur = cur.Next
		} else {
			cur.Next = l2
			l2 = l2.Next
			cur = cur.Next
		}
	}
	if l1 != nil {
		cur.Next = l1
	} else {
		cur.Next = l2
	}
	return tmp.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	mid := getMid(head)
	left := sortList(head)
	right := sortList(mid)
	return merge(left, right)
}

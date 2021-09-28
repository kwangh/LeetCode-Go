package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}

	cur := head
	i := 0
	for cur != nil && i < k {
		cur = cur.Next
		i++
	}

	if i != k {
		return head
	}

	last, cur := reverseKGroup(cur, k), head
	for i > 0 {
		i--
		tmp := cur
		cur = cur.Next
		tmp.Next = last
		last = tmp
	}
	return last
}

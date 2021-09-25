package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	tmp := head
	l := 1
	for tmp.Next != nil {
		l++
		tmp = tmp.Next
	}
	if k == 0 {
		return head
	}
	tmp.Next = head
	k = k % l
	tmp = head
	for i := 0; i < l-k-1; i++ {
		tmp = tmp.Next
	}
	res := tmp.Next
	tmp.Next = nil
	return res
}

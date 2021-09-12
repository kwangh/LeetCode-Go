package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	for head != nil && head.Next != nil {
		a := head
		b := head.Next

		prev.Next = b
		a.Next = b.Next
		b.Next = a

		prev = a
		head = head.Next
	}

	return dummy.Next
}

func swapPairsRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	first := head.Next
	head.Next = swapPairs(head.Next.Next)
	first.Next = head

	return first
}

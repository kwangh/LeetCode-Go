package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	t, b := head, head
	for b != nil && b.Next != nil {
		b = b.Next.Next
		t = t.Next
		if t == b {
			return true
		}
	}
	return false
}

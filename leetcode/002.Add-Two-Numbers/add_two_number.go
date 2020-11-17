package leetcode

// ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers function adds two list numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	cur := res

	l1 = &ListNode{Next: l1}
	l2 = &ListNode{Next: l2}

	for *l1.Next != (ListNode{}) {
		l1 = l1.Next
		l2 = l2.Next

		sum := cur.Val + l1.Val + l2.Val
		mod := sum % 10

		cur.Val = mod
		cur.Next = &ListNode{}
		cur = cur.Next
		cur.Val = (sum - mod) / 10
	}

	return res
}

package leetcode

// ListNode structure
type ListNode struct {
	Val  int
	Next *ListNode
}

// AddTwoNumbers function adds two list numbers
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	res := &ListNode{}
	n1, n2, carry, sum, cur := 0, 0, 0, 0, res
	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		} else {
			n1 = 0
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		} else {
			n2 = 0
		}
		sum = n1 + n2 + carry
		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
		carry = sum / 10
	}

	return res.Next
}

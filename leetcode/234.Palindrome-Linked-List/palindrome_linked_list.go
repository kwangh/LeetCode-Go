package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	var rev *ListNode
	for slow != nil {
		tmp := slow.Next
		slow.Next = rev
		rev = slow
		slow = tmp
	}
	for rev != nil {
		if rev.Val != head.Val {
			return false
		}
		rev = rev.Next
		head = head.Next
	}

	return true
}

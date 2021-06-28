package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fast := head
	var rev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head.Next, rev, head = rev, head, head.Next
	}
	if fast != nil {
		head = head.Next
	}
	for rev != nil {
		if rev.Val != head.Val {
			return false
		}
		rev, head = rev.Next, head.Next
	}
	return true
}

func isPalindromeRestore(head *ListNode) bool {
	fast := head
	var rev *ListNode
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		head.Next, rev, head = rev, head, head.Next
	}

	var tail *ListNode
	if fast != nil {
		tail = head.Next
	} else {
		tail = head
	}

	res := true
	for rev != nil {
		res = res && rev.Val == tail.Val
		rev.Next, head, rev = head, rev, rev.Next
		tail = tail.Next
	}
	return res
}

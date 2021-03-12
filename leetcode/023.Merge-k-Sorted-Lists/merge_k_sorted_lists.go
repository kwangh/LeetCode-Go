package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return &ListNode{}
	}
	if len(lists) == 1 {
		return lists[0]
	}

	res := lists[0]
	for i := 1; i < len(lists)-1; i++ {
		a := lists[i]
		t := res
		for a != nil && t != nil {
			if a.Val < t.Val {
				tmp := &ListNode{Val: a.Val, Next: t}
				a = a.Next
			} else {
				t = t.Next
			}
		}
		for a != nil {
			res.Next = a
			a = a.Next
		}
	}
	return res
}

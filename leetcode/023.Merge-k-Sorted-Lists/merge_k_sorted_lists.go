package leetcode

import (
	"sort"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	s := []int{}
	for i := 0; i < len(lists); i++ {
		list := lists[i]
		for list != nil {
			s = append(s, list.Val)
			list = list.Next
		}
	}

	res := &ListNode{}
	t := res
	sort.Ints(s)
	for i := 0; i < len(s); i++ {
		t.Next = &ListNode{Val: s[i]}
		t = t.Next
	}

	return res.Next
}

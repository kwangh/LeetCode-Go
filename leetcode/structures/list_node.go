package structures

import "fmt"

// ListNode list node
type ListNode struct {
	Val  int
	Next *ListNode
}

// Ints returns int slices of list node
func Ints(head *ListNode) []int {
	limit, times := 100, 0
	res := []int{}
	for head != nil {
		times++
		if times > limit {
			fmt.Println("list nodes contain more than 100 elements")
			break
		}
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}

// Lists returns list nodes
func Lists(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}

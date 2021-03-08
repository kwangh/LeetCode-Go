package structures

import "fmt"

// ListNodes list node
type ListNodes struct {
	Val  int
	Next *ListNodes
}

// Ints returns int slices of list node
func Ints(head *ListNodes) []int {
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
func Lists(nums []int) *ListNodes {
	if len(nums) == 0 {
		return nil
	}
	l := &ListNodes{}
	t := l
	for _, v := range nums {
		t.Next = &ListNodes{Val: v}
		t = t.Next
	}
	return l.Next
}

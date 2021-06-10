package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetIntersectionNode(t *testing.T) {
	a := assert.New(t)
	c1 := &ListNode{Val: 8, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}
	a1 := &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: c1}}
	b1 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 1, Next: c1}}}
	a.Equal(c1, getIntersectionNode(a1, b1))
	c1 = &ListNode{Val: 2, Next: &ListNode{Val: 4}}
	a1 = &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: &ListNode{Val: 1, Next: c1}}}
	b1 = &ListNode{Val: 3, Next: c1}
	a.Equal(c1, getIntersectionNode(a1, b1))
	b1 = &ListNode{Val: 1, Next: &ListNode{Val: 5}}
	a1 = &ListNode{Val: 2, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	a.Equal((*ListNode)(nil), getIntersectionNode(a1, b1))
}

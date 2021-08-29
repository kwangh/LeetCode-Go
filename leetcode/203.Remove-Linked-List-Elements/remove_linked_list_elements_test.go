package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElements(t *testing.T) {
	a := assert.New(t)
	a.Equal(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, removeElements(&ListNode{1, &ListNode{2, &ListNode{6, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{6, nil}}}}}}}, 6))
	var v *ListNode
	a.Equal(v, removeElements(v, 1))
	a.Equal(v, removeElements(&ListNode{7, &ListNode{7, &ListNode{7, &ListNode{7, nil}}}}, 7))
}

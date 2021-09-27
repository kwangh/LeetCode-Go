package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseKGroup(t *testing.T) {
	a := assert.New(t)
	a.Equal(&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, &ListNode{5, nil}}}}}, reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	a.Equal(&ListNode{3, &ListNode{2, &ListNode{1, &ListNode{4, &ListNode{5, nil}}}}}, reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 3))
	a.Equal(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, reverseKGroup(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 1))
	a.Equal(&ListNode{1, nil}, reverseKGroup(&ListNode{1, nil}, 1))

}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateRight(t *testing.T) {
	a := assert.New(t)
	a.Equal(&ListNode{4, &ListNode{5, &ListNode{1, &ListNode{2, &ListNode{3, nil}}}}}, rotateRight(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}, 2))
	a.Equal(&ListNode{2, &ListNode{0, &ListNode{1, nil}}}, rotateRight(&ListNode{0, &ListNode{1, &ListNode{2, nil}}}, 4))
	a.Equal(&ListNode{1, nil}, rotateRight(&ListNode{1, nil}, 1))
}

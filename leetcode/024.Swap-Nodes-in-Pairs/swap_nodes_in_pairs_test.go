package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwapPairs(t *testing.T) {
	a := assert.New(t)
	a.Equal(&ListNode{}, swapPairs(&ListNode{}))
	a.Equal(&ListNode{2, &ListNode{1, &ListNode{4, &ListNode{3, nil}}}}, swapPairs(&ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, nil}}}}))
	a.Equal(&ListNode{Val: 1}, swapPairs(&ListNode{Val: 1}))
}

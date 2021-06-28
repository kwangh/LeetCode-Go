package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}))
	a.Equal(false, isPalindrome(&ListNode{Val: 1, Next: &ListNode{Val: 2}}))
	a.Equal(true, isPalindrome(&ListNode{Val: 1}))
}

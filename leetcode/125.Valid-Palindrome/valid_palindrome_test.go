package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isPalindrome("A man, a plan, a canal: Panama"))
	a.Equal(false, isPalindrome("race a car"))
	a.Equal(true, isPalindrome(".,"))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isPalindrome(121), "isPalindrome(121) should be true")
	a.Equal(false, isPalindrome(-121), "isPalindrome(-121) should be false")
	a.Equal(false, isPalindrome(10), "isPalindrome(10) should be false")
	a.Equal(false, isPalindrome(-101), "isPalindrome(-101) should be false")
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, lengthOfLongestSubstring("abcabcbb"))
	a.Equal(1, lengthOfLongestSubstring("bbbbb"))
	a.Equal(3, lengthOfLongestSubstring("pwwkew"))
	a.Equal(0, lengthOfLongestSubstring(""))
}

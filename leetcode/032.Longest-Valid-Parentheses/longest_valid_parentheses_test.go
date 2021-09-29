package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestValidParentheses(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, longestValidParentheses("(()"))
	a.Equal(2, longestValidParentheses("())"))
	a.Equal(4, longestValidParentheses(")()())"))
	a.Equal(0, longestValidParentheses(""))
}

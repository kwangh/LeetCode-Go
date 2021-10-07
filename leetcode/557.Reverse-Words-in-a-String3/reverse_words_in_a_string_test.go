package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	a := assert.New(t)
	a.Equal("s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
	a.Equal("doG gniD", reverseWords("God Ding"))
}

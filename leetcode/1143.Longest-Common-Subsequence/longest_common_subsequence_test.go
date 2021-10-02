package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, longestCommonSubsequence("abcde", "ace"))
}

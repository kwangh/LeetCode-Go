package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWord(t *testing.T) {
	a := assert.New(t)
	a.Equal(5, lengthOfLastWord("Hello World"))
	a.Equal(0, lengthOfLastWord(" "))
	a.Equal(1, lengthOfLastWord("a "))
}

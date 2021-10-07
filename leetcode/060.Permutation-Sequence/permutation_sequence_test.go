package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPermutation(t *testing.T) {
	a := assert.New(t)
	a.Equal("213", getPermutation(3, 3))
	a.Equal("2314", getPermutation(4, 9))
	a.Equal("123", getPermutation(3, 1))
}

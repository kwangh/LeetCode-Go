package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	a := assert.New(t)
	b := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
	reverseString(b)
	a.Equal([]byte{'h', 'a', 'n', 'n', 'a', 'H'}, b)
	b = []byte{'h', 'e', 'l', 'l', 'o'}
	reverseString(b)
	a.Equal([]byte{'o', 'l', 'l', 'e', 'h'}, b)
}

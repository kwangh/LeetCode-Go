package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeBitwiseAnd(t *testing.T) {
	a := assert.New(t)
	a.Equal(4, rangeBitwiseAnd(5, 7))
	a.Equal(0, rangeBitwiseAnd(0, 0))
	a.Equal(0, rangeBitwiseAnd(1, 2147483647))
}

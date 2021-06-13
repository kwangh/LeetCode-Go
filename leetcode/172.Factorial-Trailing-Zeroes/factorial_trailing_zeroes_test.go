package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrailingZeroes(t *testing.T) {
	a := assert.New(t)
	a.Equal(0, trailingZeroes(3), "3! = 6, no trailing zero.")
	a.Equal(1, trailingZeroes(5), "5! = 120, one trailing zero.")
	a.Equal(0, trailingZeroes(0), "0")
	a.Equal(7, trailingZeroes(30))
}

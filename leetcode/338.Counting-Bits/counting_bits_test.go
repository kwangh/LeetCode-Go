package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{0, 1, 1}, countBits(2))
	a.Equal([]int{0, 1, 1, 2, 1, 2}, countBits(5))
}

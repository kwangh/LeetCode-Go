package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, jump([]int{2, 3, 1, 1, 4}))
	a.Equal(2, jump([]int{2, 3, 0, 1, 4}))
	a.Equal(1, jump([]int{2, 1}))
	a.Equal(1, jump([]int{1, 2}))
	a.Equal(0, jump([]int{0}))
}

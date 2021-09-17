package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, jump([]int{2, 3, 1, 1, 4}))
	a.Equal(2, jump([]int{2, 3, 0, 1, 4}))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveZeroes(t *testing.T) {
	a := assert.New(t)
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	a.Equal([]int{1, 3, 12, 0, 0}, nums)
	nums = []int{0}
	moveZeroes(nums)
	a.Equal([]int{0}, nums)
	nums = []int{0, 1, 0, 3, 0, 12}
	moveZeroes(nums)
	a.Equal([]int{1, 3, 12, 0, 0, 0}, nums)
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchInsert(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, searchInsert([]int{1, 3, 5, 6}, 5))
	a.Equal(1, searchInsert([]int{1, 3, 5, 6}, 2))
	a.Equal(4, searchInsert([]int{1, 3, 5, 6}, 7))
	a.Equal(0, searchInsert([]int{1, 3, 5, 6}, 0))
	a.Equal(0, searchInsert([]int{1}, 0))
}

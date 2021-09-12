package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFourSum(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}, fourSum([]int{1, 0, -1, 0, -2, 2}, 0))
	a.Equal([][]int{{2, 2, 2, 2}}, fourSum([]int{2, 2, 2, 2, 2}, 8))
	a.Equal([][]int{{2, 2, 2, 2}}, fourSum([]int{2, 2, 2, 2, 2, 2}, 8))
	a.Equal([][]int{{0, 0, 0, 0}}, fourSum([]int{0, 0, 0, 0}, 0))
}

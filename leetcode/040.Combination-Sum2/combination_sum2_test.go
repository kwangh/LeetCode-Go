package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum2(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}

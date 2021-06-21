package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDifference(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{2, 1, 4, 1}, minDifference([]int{1, 3, 4, 8}, [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 3}}))
	a.Equal([]int{-1, 1, 1, 3}, minDifference([]int{4, 5, 2, 2, 7, 10}, [][]int{{2, 3}, {0, 2}, {0, 5}, {3, 5}}))
}

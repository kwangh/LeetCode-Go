package combinationsum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCombinationSum(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{2, 2, 3}, {7}}, combinationSum([]int{2, 3, 6, 7}, 7))
	a.Equal([][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, combinationSum([]int{2, 3, 5}, 8))
	a.Equal([][]int{}, combinationSum([]int{2}, 1))
	a.Equal([][]int{{1}}, combinationSum([]int{1}, 1))
	a.Equal([][]int{{1, 1}}, combinationSum([]int{1}, 2))
}

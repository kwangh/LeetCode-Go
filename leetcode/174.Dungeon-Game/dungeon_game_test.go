package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateMinimumHP(t *testing.T) {
	a := assert.New(t)
	a.Equal(7, calculateMinimumHP([][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}))
	a.Equal(1, calculateMinimumHP([][]int{{0}}))
}

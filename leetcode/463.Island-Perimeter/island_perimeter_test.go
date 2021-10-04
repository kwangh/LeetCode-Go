package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIslandPerimeter(t *testing.T) {
	a := assert.New(t)
	a.Equal(16, islandPerimeter([][]int{{0, 1, 0, 0}, {1, 1, 1, 0}, {0, 1, 0, 0}, {1, 1, 0, 0}}))
	a.Equal(4, islandPerimeter([][]int{{1}}))
	a.Equal(4, islandPerimeter([][]int{{1, 0}}))
}

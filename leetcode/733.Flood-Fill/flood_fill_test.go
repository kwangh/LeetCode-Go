package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFloodFill(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}, floodFill([][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2))
	a.Equal([][]int{{0, 0, 0}, {0, 1, 1}}, floodFill([][]int{{0, 0, 0}, {0, 1, 1}}, 1, 1, 1))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridGame(t *testing.T) {
	a := assert.New(t)
	a.Equal(int64(4), gridGame([][]int{{2, 5, 4}, {1, 5, 1}}))
	a.Equal(int64(4), gridGame([][]int{{3, 3, 1}, {8, 5, 2}}))
	a.Equal(int64(7), gridGame([][]int{{1, 3, 1, 15}, {1, 3, 3, 1}}))
}

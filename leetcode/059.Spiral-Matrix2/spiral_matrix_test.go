package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMatrix(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}, generateMatrix(3))
	a.Equal([][]int{{1}}, generateMatrix(1))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTotalNQueens(t *testing.T) {
	a := assert.New(t)
	a.Equal(2, totalNQueens(4))
	a.Equal(1, totalNQueens(1))
}

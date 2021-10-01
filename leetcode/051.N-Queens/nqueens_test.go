package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSolveNQueens(t *testing.T) {
	a := assert.New(t)
	a.Equal([][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}, solveNQueens(4))
	a.Equal([][]string{{"Q"}}, solveNQueens(1))
}

package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPreorderTraversal(t *testing.T) {
	a := assert.New(t)
	a.Equal([]int{1, 2, 3}, preorderTraversal(&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}))
}

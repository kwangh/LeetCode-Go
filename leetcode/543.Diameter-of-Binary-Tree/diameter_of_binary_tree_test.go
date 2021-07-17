package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	a := assert.New(t)
	a.Equal(3, diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}))
	a.Equal(1, diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}))
}

package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}))
	a.Equal(false, isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}}, &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}))
	a.Equal(false, isSameTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 1}}, &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}))
}

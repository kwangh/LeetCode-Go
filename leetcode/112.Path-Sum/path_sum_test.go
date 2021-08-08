package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHasPathSum(t *testing.T) {
	a := assert.New(t)
	a.Equal(true, hasPathSum(&TreeNode{Val: 5,
		Left: &TreeNode{Val: 4,
			Left: &TreeNode{Val: 11,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 2},
			},
			Right: nil},
		Right: &TreeNode{Val: 8,
			Left: &TreeNode{Val: 13},
			Right: &TreeNode{Val: 4,
				Left:  nil,
				Right: &TreeNode{Val: 1}}}}, 22))
	a.Equal(false, hasPathSum(&TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3}}, 5))
	a.Equal(false, hasPathSum(&TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2},
		Right: nil}, 0))
	a.Equal(false, hasPathSum(nil, 0))
}

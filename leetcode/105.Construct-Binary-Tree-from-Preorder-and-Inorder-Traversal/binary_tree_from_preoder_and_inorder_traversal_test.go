package leetcode

import (
	"fmt"
	"testing"
)

func TestBuildTree(t *testing.T) {
	cases := []struct {
		preorder []int
		inorder  []int
		want     *TreeNode
	}{
		{
			[]int{1, 2, 4, 5, 3, 6, 7},
			[]int{4, 2, 5, 1, 6, 3, 7},
			&TreeNode{Val: 1,
				Left: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}},
				Right: &TreeNode{Val: 3,
					Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}},
		},
		/*{
			[]int{3, 9, 20, 15, 7},
			[]int{9, 3, 15, 20, 7},
			&TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		},*/
	}
	for _, c := range cases {
		got := buildTree(c.preorder, c.inorder)
		fmt.Printf("want:%#v\ngot:%#v\n", c.want, got)
	}
}

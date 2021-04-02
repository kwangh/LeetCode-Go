package leetcode

import "testing"

func TestIsValidBST(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want bool
	}{
		{
			&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			true,
		},
		{
			&TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}},
			false,
		},
		{
			&TreeNode{Val: 0},
			true,
		},
		{
			&TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 7}}},
			false,
		},
	}
	for _, c := range cases {
		got := isValidBST(c.root)
		if c.want != got {
			t.Errorf("want:%t instead got:%t", c.want, got)
		}
	}
}

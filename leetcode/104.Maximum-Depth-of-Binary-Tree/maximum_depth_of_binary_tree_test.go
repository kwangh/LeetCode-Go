package leetcode

import "testing"

func TestMaxDepth(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want int
	}{
		{&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}, 3},
		{&TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, 2},
		{nil, 0},
		{&TreeNode{Val: 0}, 1},
	}
	for _, c := range cases {
		got := maxDepth(c.root)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

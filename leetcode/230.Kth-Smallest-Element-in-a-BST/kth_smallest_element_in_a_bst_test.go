package leetcode

import "testing"

func TestKthSmallest(t *testing.T) {
	cases := []struct {
		root *TreeNode
		k    int
		want int
	}{
		{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}, k: 1, want: 1},
		{root: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6}}, k: 3, want: 3},
	}
	for _, c := range cases {
		got := kthSmallest(c.root, c.k)
		if c.want != got {
			t.Errorf("want:%d instead got:%d", c.want, got)
		}
	}
}

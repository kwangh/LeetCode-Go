package leetcode

import "testing"

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestInorderTraversal(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want []int
	}{
		{
			&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}},
			[]int{1, 3, 2},
		},
	}
	for _, c := range cases {
		got := inorderTraversal(c.root)
		if !equal(c.want, got) {
			t.Errorf("want:%v instead got:%v", c.want, got)
		}
	}
}

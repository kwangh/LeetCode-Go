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

func TestLevelOrder(t *testing.T) {
	cases := []struct {
		root *TreeNode
		want [][]int
	}{
		{
			nil, [][]int{},
		},
		{
			&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			[][]int{{3}, {9, 20}, {15, 7}},
		},
	}
	for _, c := range cases {
		got := levelOrder(c.root)
		if len(c.want) != len(got) {
			t.Errorf("want:%v instead got:%v", c.want, got)
			break
		}
		for i, row := range c.want {
			if !equal(row, got[i]) {
				t.Errorf("want:%v instead got:%v", c.want, got)
			}
		}
	}
}

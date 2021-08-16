package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var s []*TreeNode
	for root != nil || len(s) != 0 {
		for root != nil {
			s = append(s, root)
			res = append(res, root.Val)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		root = root.Right
	}
	return res
}

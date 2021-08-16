package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var s []*TreeNode
	for root != nil || len(s) != 0 {
		for root != nil {
			s = append(s, root)
			res = append(res, 0)
			copy(res[1:], res)
			res[0] = root.Val
			root = root.Right
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		root = root.Left
	}
	return res
}

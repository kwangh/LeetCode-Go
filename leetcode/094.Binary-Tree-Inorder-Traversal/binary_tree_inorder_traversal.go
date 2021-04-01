package leetcode

// TreeNode tree node
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	traverse(root.Left, res)
	*res = append(*res, root.Val)
	traverse(root.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
	var res []int
	traverse(root, &res)
	return res
}

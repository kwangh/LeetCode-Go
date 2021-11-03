package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, val int) int {
	if root == nil {
		return 0
	}
	cur := val*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return cur
	}
	return dfs(root.Left, cur) + dfs(root.Right, cur)
}

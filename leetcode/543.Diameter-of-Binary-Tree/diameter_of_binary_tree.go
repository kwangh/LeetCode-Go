package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recursive(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	l, r := recursive(root.Left, max), recursive(root.Right, max)
	if l+r > *max {
		*max = l + r
	}
	if l > r {
		return l + 1
	}
	return r + 1
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max int
	recursive(root, &max)
	return max
}

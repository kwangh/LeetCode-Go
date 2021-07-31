package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := height(root.Left), height(root.Right)
	if left == -1 || right == -1 {
		return -1
	}
	diff := left - right
	if diff < 0 {
		diff = -diff
	}
	if diff > 1 {
		return -1
	}
	return max(left, right) + 1
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

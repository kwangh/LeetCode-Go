package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func valid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if (min != nil && root.Val <= min.Val) || (max != nil && root.Val >= max.Val) {
		return false
	}
    
	return valid(root.Left, min, root) && valid(root.Right, root, max)
}

func isValidBST(root *TreeNode) bool {
	return valid(root, nil, nil)
}

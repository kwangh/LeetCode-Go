package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	return pathSum(root.Left, targetSum) + pathSum(root.Right, targetSum) + pathSumFrom(root, targetSum)
}

func pathSumFrom(root *TreeNode, target int) int {
	if root == nil {
		return 0
	}

	var res int
	if root.Val == target {
		res++
	}
	return res + pathSumFrom(root.Left, target-root.Val) + pathSumFrom(root.Right, target-root.Val)
}

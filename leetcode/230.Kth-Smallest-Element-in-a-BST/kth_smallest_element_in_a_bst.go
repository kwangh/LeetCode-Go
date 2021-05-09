package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	var s []*TreeNode
	for root != nil || len(s) != 0 {
		for root != nil {
			s = append(s, root)
			root = root.Left
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		k--
		if k == 0 {
			break
		}
		root = root.Right
	}
	return root.Val
}

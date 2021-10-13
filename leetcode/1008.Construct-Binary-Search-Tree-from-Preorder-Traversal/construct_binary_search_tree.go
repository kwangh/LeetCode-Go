package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var i int
	for i = 1; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			break
		}
	}
	root := &TreeNode{preorder[0], nil, nil}
	root.Left = bstFromPreorder(preorder[1:i])
	root.Right = bstFromPreorder(preorder[i:])
	return root
}

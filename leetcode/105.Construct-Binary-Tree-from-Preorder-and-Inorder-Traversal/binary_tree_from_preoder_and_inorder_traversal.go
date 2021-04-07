package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var preorderIndex int
var m map[int]int

func arrayToTree(preorder []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	rootVal := preorder[preorderIndex]
	preorderIndex++
	root := &TreeNode{Val: rootVal}
	root.Left = arrayToTree(preorder, left, m[rootVal]-1)
	root.Right = arrayToTree(preorder, m[rootVal]+1, right)
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	preorderIndex = 0
	m = make(map[int]int)
	for i, v := range inorder {
		m[v] = i
	}
	return arrayToTree(preorder, 0, len(preorder)-1)
}

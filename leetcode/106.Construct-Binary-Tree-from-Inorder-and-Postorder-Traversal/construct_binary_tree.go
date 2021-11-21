package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var index map[int]int

func buildTree(inorder []int, postorder []int) *TreeNode {
	index = make(map[int]int, len(inorder))
	for i, v := range inorder {
		index[v] = i
	}
	n := len(inorder) - 1
	return dfs(inorder, postorder, 0, n, 0, n)
}

func dfs(io []int, po []int, is, ie, ps, pe int) *TreeNode {
	if is > ie {
		return nil
	}

	rootIndex := index[po[pe]]
	pe--
	root := &TreeNode{Val: io[rootIndex]}
	root.Left = dfs(io, po, is, rootIndex-1, ps, ps+rootIndex-1-is)
	root.Right = dfs(io, po, rootIndex+1, ie, ps+rootIndex-is, pe)
	return root
}

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	for root != nil || len(queue) > 0 {
		var cur []int
		var next []*TreeNode
		for len(queue) > 0 {
			root = queue[0]
			queue = queue[1:]
			if root != nil {
				cur = append(cur, root.Val)
				next = append(next, root.Left, root.Right)
			}
		}
		if cur != nil {
			res = append(res, cur)
			queue = append(queue, next...)
		}
	}
	return res
}

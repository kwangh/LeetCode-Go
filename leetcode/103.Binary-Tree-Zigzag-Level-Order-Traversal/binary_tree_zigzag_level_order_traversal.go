package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var cur []int
		var next []*TreeNode
		for len(queue) > 0 {
			root = queue[0]
			queue = queue[1:]

			if root != nil {
				cur = append(cur, root.Val)
				if root.Left != nil {
					next = append(next, root.Left)
				}
				if root.Right != nil {
					next = append(next, root.Right)
				}
			}
		}
		if cur != nil {
			res = append(res, cur)
			queue = append(queue, next...)
		}
	}
	for i, row := range res {
		if i%2 == 1 {
			for j := 0; j < len(row)/2; j++ {
				row[j], row[len(row)-1-j] = row[len(row)-1-j], row[j]
			}
		}
	}
	return res
}

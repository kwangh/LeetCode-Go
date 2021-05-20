package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func symmetric(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if (l == nil && r != nil) || (l != nil && r == nil) {
		return false
	}

	if l.Val == r.Val {
		return symmetric(l.Left, r.Right) && symmetric(l.Right, r.Left)
	} else {
		return false
	}
}

func isSymmetric(root *TreeNode) bool {
	return symmetric(root.Left, root.Right)
}

func isSymmetricIteratively(root *TreeNode) bool {
    var l, r []*TreeNode
	leftRoot, rightRoot := root.Left, root.Right
	for leftRoot != nil || rightRoot != nil || len(l) != 0 || len(r) != 0 {
		for leftRoot != nil && rightRoot != nil {
			if leftRoot.Val != rightRoot.Val {
				return false
			}
			l = append(l, leftRoot)
			leftRoot = leftRoot.Left
			r = append(r, rightRoot)
			rightRoot = rightRoot.Right
		}
        if (leftRoot == nil && rightRoot != nil) || (leftRoot != nil && rightRoot == nil) {
			return false
		}
		leftRoot = l[len(l)-1]
		l = l[:len(l)-1]
		rightRoot = r[len(r)-1]
		r = r[:len(r)-1]
		if leftRoot.Val != rightRoot.Val {
			return false
		}
		leftRoot = leftRoot.Right
		rightRoot = rightRoot.Left
	}
	return true
}
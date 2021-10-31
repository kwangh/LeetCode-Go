package leetcode

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	cur := root
	for cur != nil {
		if cur.Child == nil {
			cur = cur.Next
			continue
		}

		tmp := cur.Child
		for tmp.Next != nil {
			tmp = tmp.Next
		}
		tmp.Next = cur.Next
		if cur.Next != nil {
			cur.Next.Prev = tmp
		}
		cur.Next = cur.Child
		cur.Child.Prev = cur
		cur.Child = nil
	}
	return root
}

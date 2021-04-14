package leetcode

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	cur := head
	for cur != nil {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next.Next
	}

	cur = head
	for cur != nil {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}

	res := &Node{}
	copy := res
	cur = head
	for cur != nil {
		copy.Next = cur.Next
		copy = copy.Next

		cur.Next = cur.Next.Next
		cur = cur.Next
	}

	return res.Next
}

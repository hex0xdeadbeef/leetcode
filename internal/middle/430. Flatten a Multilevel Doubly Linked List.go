package middle

type NodeC struct {
	Val   int
	Prev  *NodeC
	Next  *NodeC
	Child *NodeC
}

func flattenLinkedList(root *NodeC) *NodeC {
	cur := root
	nodes := []*NodeC{}

	for ; cur != nil; cur = cur.Next {
		nodes = append(nodes, cur)
		if cur.Child != nil {
			traverseChild(cur.Child, &nodes)
		}
	}

	for i := range len(nodes) - 1 {
		nodes[i].Next, nodes[i].Child = nodes[i+1], nil
	}

	for i := len(nodes)-1; i >= 1; i-- {
		nodes[i].Prev = nodes[i-1]
	}

	return root
}

func traverseChild(root *NodeC, nodes *[]*NodeC) {
	for ; root != nil; root = root.Next {
		*nodes = append(*nodes, root)
		if root.Child != nil {
			traverseChild(root.Child, nodes)
		}
	}

}

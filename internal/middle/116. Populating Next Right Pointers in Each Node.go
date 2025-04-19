package middle

type CustomNodeTree struct {
	Val   int
	Left  *CustomNodeTree
	Right *CustomNodeTree
	Next  *CustomNodeTree
}

func connect(root *CustomNodeTree) *CustomNodeTree {
	m := map[int][]*CustomNodeTree{}

	var traverse func(*CustomNodeTree, map[int][]*CustomNodeTree, int)

	traverse = func(root *CustomNodeTree, m map[int][]*CustomNodeTree, d int) {
		if root == nil {
			return
		}

		d++

		m[d] = append(m[d], root)
		traverse(root.Left, m, d)
		traverse(root.Right, m, d)
	}

	traverse(root, m, 0)

	for _, v := range m {
		for i := range len(v) - 1 {
			v[i].Next = v[i+1]
		}
	}

	return root
}

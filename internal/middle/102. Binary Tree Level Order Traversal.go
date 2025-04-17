package middle

import . "leetcode/pkg/datastructures/treenode"

func levelOrder(root *TreeNode) [][]int {
	m := make(map[int][]int)

	var traverse func(*TreeNode, int)

	traverse = func(root *TreeNode, d int) {
		if root == nil {
			return
		}

		d++

		traverse(root.Left, d)
		m[d] = append(m[d], root.Val)
		traverse(root.Right, d)
	}

	traverse(root, -1)

	res := make([][]int, len(m))
	var i int
	for range m {
		res[i] = m[i]
		i++
	}

	return res
}

package middle

import . "leetcode/pkg/datastructures/treenode"

func rightSideView(root *TreeNode) []int {
	m := make(map[int]int)

	var traverse func(*TreeNode, map[int]int, int)

	traverse = func(root *TreeNode, m map[int]int, d int) {
		if root == nil {
			return
		}

		_, ok := m[d]
		if !ok {
			m[d] = root.Val
		}

		traverse(root.Right, m, d+1)
		traverse(root.Left, m, d+1)
	}

	traverse(root, m, 0)

	res := make([]int, 0, len(m))

	for i := range len(m) {
		res = append(res, m[i])
	}

	return res
}
